package proofs_sender

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"time"
)

type ProofsSender struct {
	log     *logan.Entry
	masterQ data.MasterQ
	cfg     *config.NetworkConfig
	client  *ethclient.Client
}

func NewProofsSender(log *logan.Entry, masterQ data.MasterQ, cfg *config.NetworkConfig) (*ProofsSender, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial connect")
	}

	return &ProofsSender{
		log:     log,
		masterQ: masterQ,
		cfg:     cfg,
		client:  client,
	}, nil
}

func (s *ProofsSender) Run(ctx context.Context) {
	s.log.Info("Starting proofs-sender")
	running.WithBackOff(
		ctx, s.log, "proofs-sender",
		s.run,
		time.Second,
		time.Second*5,
		time.Minute,
	)
}

func (s *ProofsSender) run(ctx context.Context) error {
	proofs, err := s.masterQ.Proofs().ResetFilter().OrderBy("created_at", "asc").Select()
	if err != nil {
		return errors.Wrap(err, "failed to select proofs")
	}

	for _, proof := range proofs {
		if err := s.masterQ.Transaction(func(db data.MasterQ) error {
			if err := db.Proofs().DeleteByID(proof.ID); err != nil {
				return errors.Wrap(err, "failed to delete proof by id")
			}

			dataBytes, err := hexutil.Decode(proof.TxData)
			if err != nil {
				return errors.Wrap(err, "failed to decode data")
			}

			gasPrice, err := s.client.SuggestGasPrice(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to suggest gas price")
			}

			gas, err := s.client.EstimateGas(ctx, ethereum.CallMsg{
				From:     crypto.PubkeyToAddress(s.cfg.PrivateKey.PublicKey),
				To:       &s.cfg.VerifierAddress,
				GasPrice: gasPrice,
				Data:     dataBytes,
			})
			if err != nil {
				return errors.Wrap(err, "failed to estimate gas")
			}

			s.cfg.LockNonce()
			defer s.cfg.UnlockNonce()

			tx, err := types.SignNewTx(
				s.cfg.PrivateKey,
				types.NewCancunSigner(s.cfg.ChainID),
				&types.LegacyTx{
					Nonce:    s.cfg.Nonce(),
					Gas:      gas,
					GasPrice: gasPrice,
					To:       &s.cfg.VerifierAddress,
					Data:     dataBytes,
				},
			)
			if err != nil {
				return errors.Wrap(err, "failed to sign transaction")
			}

			if err := s.client.SendTransaction(ctx, tx); err != nil {
				return errors.Wrap(err, "failed to send transaction")
			}

			s.cfg.IncrementNonce()

			return nil
		}); err != nil {
			return errors.Wrap(err, "failed to perform SQL transaction")
		}
	}

	return nil
}
