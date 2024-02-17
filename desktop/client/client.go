package client

import (
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client/api"
)

func NewClient() *api.ClientWithResponses {
	cl, err := api.NewClientWithResponses(build.APIBaseURL)
	if err != nil {
		panic(err)
	}
	return cl
}
