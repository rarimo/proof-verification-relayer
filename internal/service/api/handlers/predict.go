package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func PredictHandlers(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	var data []byte

	countTx, err := checker.GetCountTx(Config(r), data, address)

	if err != nil {
		Log(r).Errorf("Failed get count tx for wallet: %v", err)
		ape.RenderErr(w, problems.InternalError())

		return
	}

	type response struct {
		Count uint64
	}

	ape.Render(w, &response{Count: countTx})

}
