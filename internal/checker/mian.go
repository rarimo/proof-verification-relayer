package checker

import (
	"context"
	"database/sql"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	usdcabi "github.com/rarimo/proof-verification-relayer/internal/checker/usdc_abi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"

	"gitlab.com/distributed_lab/logan/v3"
)

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.

func CheckEvents(ctx context.Context, cfg config.Config) {

	pgDB := pg.NewMaterDB(cfg.DB())
	logger := cfg.Log()
	csparam := cfg.NewCheckerCfg()

	client, err := ethclient.Dial(csparam.RpcURL)

	if err != nil {
		logger.Errorf("Failed create client: %v", err)
		return
	}

	defer client.Close()

	contractAddress := common.HexToAddress(csparam.ContractAddr)

	startBlock, err := getStartBlockNumber(pgDB, logger, csparam.StartBlock)

	if err != nil {
		logger.Errorf("Failed get start block: %v", err)
		return
	}
	logger.Infof("Start block: %d", startBlock)

	if startBlock == 0 {
		startBlock, err = client.BlockNumber(context.Background())

		if err != nil {
			logger.Errorf("Failed get start block: %v", err)
			return
		}
	}

	fromAddresses := []common.Address{}
	toAddresses := []common.Address{}
	for {

		select {
		case <-ctx.Done():
			logger.Warn("unsubscribe from events")
			return
		default:

		}

		latestBlock, err := client.BlockNumber(context.Background())

		if err != nil {
			logger.Warnf("Failed to get current block: %v", err)

			select {
			case <-ctx.Done():
				logger.Warn("unsubscribe from events")
				return
			default:
				time.Sleep(10 * time.Second)

			}
			continue
		}

		query := bind.FilterOpts{
			Start:   startBlock,
			End:     &latestBlock,
			Context: context.Background(),
		}
		// Instead of "abiusdc", use the ABI package
		// that will be generated for the required contract.
		contract, err := usdcabi.NewUsdcAbiFilterer(contractAddress, client)

		if err != nil {
			logger.Fatalf("Failed to create filterer: %v", err)
		}
		logs, err := contract.FilterTransfer(&query, fromAddresses, toAddresses)

		if err != nil {
			logger.Warnf("Failed to get logs: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for logs.Next() {
			event := logs.Event

			processLog(event, pgDB, logger)
		}

		startBlock = latestBlock + 1
		err = pgDB.CheckerQ().UpdateBlock(startBlock)

		if err != nil {
			logger.Warnf("Failed update block number in db: %v", err)

		}
		select {
		case <-ctx.Done():
			logger.Info("unsubscribe from events")
			return
		default:
			time.Sleep(10 * time.Second)

		}

	}
}

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.
func processLog(event *usdcabi.UsdcAbiTransfer, pg data.CheckerDB, logger *logan.Entry) {

	to := event.To.Hex()

	var userWallet data.Wallet

	userWallet.Address = to
	userBalance, err := getUserBalance(pg, to, logger)

	if err != nil {
		logger.Errorf("Failed get user balance: %v", err)
		return
	}

	value := event.Value

	userWallet.Balance = userBalance + value.Uint64()

	err = pg.CheckerQ().UpdateBalance(userWallet)

	if err != nil {
		logger.Errorf("Failed Update user balance: %v", err)
		return
	}

}

func getUserBalance(pg data.CheckerDB, address string, logger *logan.Entry) (uint64, error) {

	userBalance, err := pg.CheckerQ().GetBalance(address)

	if err == sql.ErrNoRows {

		newWallet := &data.Wallet{
			Address: address,
			Balance: 0,
		}

		err := pg.CheckerQ().Insert(*newWallet)
		if err != nil {
			logger.Errorf("Failed insert new user wallet: %v", err)
			return 0, err
		}
		return 0, nil

	}

	return userBalance, err
}

func IsEnough(cfg config.Config, data []byte, addr string) (bool, error) {

	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())
	txFee, err := getTxFee(cfg)

	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return false, err
	}
	userBalance, err := getUserBalance(pg, addr, logger)

	if err != nil {
		logger.Errorf("Failed get user balance: %v", err)
		return false, err
	}
	if txFee <= userBalance {
		return true, nil
	}
	return false, nil

}

func GetCountTx(cfg config.Config, data []byte, addr string) (uint64, error) {

	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())
	txFee, err := getTxFee(cfg)

	if err != nil {
		logger.Errorf("Failed get count tx for ETH: %v", err)
		return 0, err
	}
	userBalance, err := getUserBalance(pg, addr, logger)

	if err != nil {
		logger.Errorf("Failed get user balance: %v", err)
		return 0, err
	}

	countTx := userBalance / txFee

	return countTx, nil
}

func getTxFee(cfg config.Config) (uint64, error) {

	logger := cfg.Log()
	csparam := cfg.NewCheckerCfg()
	pgDB := pg.NewMaterDB(cfg.DB())

	client, err := ethclient.Dial(csparam.RpcURL)

	if err != nil {
		logger.Errorf("Failed create client: %v", err)
		return 0, err
	}

	defer client.Close()

	feeCap, _ := client.SuggestGasPrice(context.Background())

	gasLimit, err := pgDB.CheckerQ().GetGasLimit()

	if err != nil {
		return 0, err
	}
	if gasLimit == 0 {

		err := pgDB.CheckerQ().InsertGasLimit(csparam.GasLimit)

		if err != nil {
			logger.Errorf("Failed insert gas limit in db: %v", err)
			return 0, err
		}

		gasLimit = csparam.GasLimit
	}

	totalFee := gasLimit * feeCap.Uint64()
	return totalFee, nil

}

func CheckUpdateGasLimit(value uint64, cfg config.Config) error {

	pgDB := pg.NewMaterDB(cfg.DB())
	logger := cfg.Log()
	csParam := cfg.NewCheckerCfg()
	gasLimit, err := pgDB.CheckerQ().GetGasLimit()

	if err != nil {
		logger.Errorf("Failed get gas limit: %v", err)
		return err
	}

	switch gasLimit {
	case 0:
		err := pgDB.CheckerQ().InsertGasLimit(csParam.GasLimit)

		if err != nil {
			logger.Errorf("Failed insert gas limit in db: %v", err)
			return err
		}
	case csParam.GasLimit:
		err := pgDB.CheckerQ().SetGasLimit(value, csParam.GasLimit)

		if err != nil {
			logger.Errorf("Failed set gas limit in db: %v", err)
			return err
		}
	default:
	}

	return nil
}

func getStartBlockNumber(pgDB data.CheckerDB, logger *logan.Entry, defaultSBlock uint64) (uint64, error) {

	block, err := pgDB.CheckerQ().GetBlock()

	if err != nil {
		logger.Errorf("Failed get block: %v", err)
		return 0, err
	}

	switch block {
	case 0:
		err := pgDB.CheckerQ().InsertBlock(defaultSBlock)
		block = defaultSBlock
		if err != nil {
			logger.Errorf("Failed insert gas limit in db: %v", err)
			return 0, err
		}
	default:

	}

	return block, nil
}
