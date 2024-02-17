package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/phimtorr/phimtor/desktop/client/api"
)

type Handler struct {
	client api.ClientWithResponsesInterface
}

func New(client api.ClientWithResponsesInterface) *Handler {
	if client == nil {
		panic("client is required")
	}
	return &Handler{
		client: client,
	}
}

func (h *Handler) Register(e *echo.Echo) {
	e.GET("/", h.Home)
	e.GET("/shows", h.ListShows)
}
