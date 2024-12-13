package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

func NewServer(cfg *Config) *http.Server {
	return &http.Server{Addr: fmt.Sprintf(":%d", cfg.Port)}
}

func StartServer(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						panic(err)
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
