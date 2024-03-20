package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"net/http"
)

func NewServer(router *chi.Mux) *http.Server {
	return &http.Server{Addr: ":8080", Handler: router}
}

func StartServer(lc fx.Lifecycle, server *http.Server, logger *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			logger.Info().Msg(fmt.Sprintf("HTTP server listening on %s", server.Addr))
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
			logger.Info().Msg("Shutting down HTTP server")
			return server.Shutdown(ctx)
		},
	})
}
