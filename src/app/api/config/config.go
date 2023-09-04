package cfg

import (
	"errors"

	"github.com/gabrielsoaressantos/env/v8"
	"github.com/rs/zerolog/log"
)

var config *Config

type Config struct {
	Server   Server   `env:"<parse>"`
	Postgres Database `env:"<parse>"`
	Mail     Mail     `env:"<parse>"`
	Strings  Strings  `env:"<parse>"`
}

func Env() *Config {
	if config == nil {
		config = &Config{}
		if err := env.ParseNested(config); err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to parse config from environment variables")
		}
	}

	return config
}

func Validate() error {
	return errors.New("implement error validation")
}
