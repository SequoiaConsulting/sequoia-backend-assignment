package model

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// loggerFor returns a logger with context related to a model
func loggerFor(model string) *zerolog.Logger {
	logger := log.With().Str("model", model).Logger()
	return &logger
}
