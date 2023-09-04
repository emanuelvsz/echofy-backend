package middlewares

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger() zerolog.Logger {
	return log.With().Str("layer", "middlewares").Logger()
}