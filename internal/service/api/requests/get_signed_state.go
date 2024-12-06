package requests

import (
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetSignedStateRequest struct {
	Root  *string `filter:"root"`
	Block *uint64 `filter:"block"`
}

func NewGetSignedStateRequest(r *http.Request) (GetSignedStateRequest, error) {
	var request GetSignedStateRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return GetSignedStateRequest{}, errors.Wrap(err, "failed to decode url")
	}

	return request, request.validate()
}

func (r GetSignedStateRequest) validate() error {
	return validation.Errors{
		"filter[root]":  validation.Validate(r.Root, validation.When(!validation.IsEmpty(validation.Match(regexp.MustCompile("^0x([0-9a-f]+)$"))))),
		"filter[block]": validation.Validate(r.Block, validation.When(!validation.IsEmpty(validation.Min(0)))),
	}.Filter()
}
