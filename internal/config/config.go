package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Transit struct {
		APIKey string
	}
}

func Load() (Config, error) {
	var c Config
	_, err := toml.DecodeFile("transit.toml", &c)
	return c, err
}
