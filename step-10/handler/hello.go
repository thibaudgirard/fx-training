package handler

import (
	"dojo-fx/step-10/repository"
	"github.com/rs/zerolog"
	"net/http"
)

type SpeakHandler struct {
	logger     *zerolog.Logger
	repository *repository.GreetRepository
}

func (s *SpeakHandler) Pattern() []Route {
	return []Route{
		{
			Path:   "/hello",
			Method: http.MethodGet,
		},
	}
}

func (s *SpeakHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Info().Msg("SpeakHandler called")
	w.Write([]byte(s.repository.Greet()))
}

func NewHelloHandler(logger *zerolog.Logger, repository *repository.GreetRepository) *SpeakHandler {
	return &SpeakHandler{
		logger:     logger,
		repository: repository,
	}
}
