package http

import (
	"dojo-fx/step-08/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(routeHandlers []handler.Handler) *chi.Mux {
	router := chi.NewRouter()

	for _, rh := range routeHandlers {
		for _, route := range rh.Pattern() {
			router.Method(route.Method, route.Path, rh)
		}
	}

	return router
}
