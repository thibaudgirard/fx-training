package main

import "github.com/go-chi/chi/v5"

func NewRouter(rh RouteHandler) *chi.Mux {
	router := chi.NewRouter()

	for _, route := range rh.Pattern() {
		router.Method(route.Method, route.Path, rh)
	}

	return router
}
