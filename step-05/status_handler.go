package main

import (
	"github.com/rs/zerolog"
	"net/http"
)

type StatusHandler struct {
	logger *zerolog.Logger
}

func NewStatusHandler(logger *zerolog.Logger) *StatusHandler {
	return &StatusHandler{
		logger: logger,
	}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info().Msg("StatusHandler called")
	w.Write([]byte("OK"))
}
