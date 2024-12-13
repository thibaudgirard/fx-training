package main

import (
	"fmt"
	"net/http"
)

func NewServer(cfg *Config) *http.Server {
	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.Port)}
}

func ServerStart(server *http.Server) {
	fmt.Printf("Server started on %s\n", server.Addr)
}
