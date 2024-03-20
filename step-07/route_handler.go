package main

import (
	"net/http"
)

type Route struct {
	Method string
	Path   string
}

type RouteHandler interface {
	http.Handler
	Pattern() []Route
}
