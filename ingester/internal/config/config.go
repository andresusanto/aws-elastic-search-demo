package config

import (
	"github.com/rs/zerolog/log"

	"github.com/kelseyhightower/envconfig"
)

// Config from env vars
type Config struct {
	Port         int16  `default:"8080" split_words:"true"`
	Region       string `default:"local" split_words:"true"`
	ESEndpoint   string `default:"http://localhost:9200" split_words:"true"`
	SignESClient bool   `default:"false" split_words:"true"`
}

// New returns the application config
func New() *Config {
	c := &Config{}
	err := envconfig.Process("jobcontroller", c)

	if err != nil {
		log.Fatal().Err(err).Msg("Environment config cannot be loaded")
	}

	return c
}
