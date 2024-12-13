package handler

import (
	"net/http"
)

type Route struct {
	Method string
	Path   string
}

type Handler interface {
	http.Handler
	Pattern() []Route
}
