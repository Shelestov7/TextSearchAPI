package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	HTTPPort     string        `envconfig:"HTTP_PORT" default:"8080"`
	FilePath     string        `envconfig:"FiLE_PATH" default:"Latin-Lipsum.txt"`
	ReadTimeout  time.Duration `envconfig:"HTTP_ADDR" default:"1s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"1s"`
}

func LoadENV() (*Config, error) {
	var c Config
	err := envconfig.Process("app", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
