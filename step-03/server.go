package main

import (
	"fmt"
	"net/http"
)

func NewServer() *http.Server {
	return &http.Server{Addr: ":8080"}
}

func ServerStart(server *http.Server) {
	fmt.Printf("Server started on %s\n", server.Addr)
}
