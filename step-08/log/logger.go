package log

import (
	"github.com/rs/zerolog"
)

func NewLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
	return &logger
}
