package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func IsEnoughHandler(w http.ResponseWriter, r *http.Request) {

	address := chi.URLParam(r, "address")

	var data []byte

	isEnough, err := checker.IsEnough(Config(r), data, address)

	if err != nil {
		Log(r).Warnf("Failed check is enough balance: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	type status struct {
		IsEnough bool
	}

	ape.Render(w, &status{IsEnough: isEnough})

}
