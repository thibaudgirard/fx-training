package main

import "net/http"

type StatusHandler struct {
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
