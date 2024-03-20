package main

import (
	"github.com/rs/zerolog"
	"net/http"
)

type HelloHandler struct {
	logger *zerolog.Logger
}

func (h *HelloHandler) Pattern() []Route {
	return []Route{
		{
			Path:   "/hello",
			Method: http.MethodGet,
		},
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info().Msg("HelloHandler called")
	w.Write([]byte("Hello world!"))
}

func NewHelloHandler(logger *zerolog.Logger) *HelloHandler {
	return &HelloHandler{
		logger: logger,
	}
}
