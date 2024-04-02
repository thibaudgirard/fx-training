package main

import "github.com/go-chi/chi/v5"

func NewRouter(sh *StatusHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Method("GET", "/status", sh)

	return router
}
