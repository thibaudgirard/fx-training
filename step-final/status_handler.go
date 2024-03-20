package main

import (
	"github.com/rs/zerolog"
	"net/http"
)

type StatusHandler struct {
	logger *zerolog.Logger
}

func (h *StatusHandler) Pattern() []Route {
	return []Route{
		{
			Path:   "/status",
			Method: http.MethodGet,
		},
	}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info().Msg("StatusHandler called")
	w.Write([]byte("OK"))
}

func NewStatusHandler(logger *zerolog.Logger) *StatusHandler {
	return &StatusHandler{
		logger: logger,
	}
}
