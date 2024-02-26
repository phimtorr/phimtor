package client

import (
	"strings"

	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client/api"
)

func NewClient() *api.ClientWithResponses {
	apiBaseURL := strings.TrimRight(build.ServerAddr, "/") + "/api/v1"
	cl, err := api.NewClientWithResponses(apiBaseURL)
	if err != nil {
		panic(err)
	}
	return cl
}
