package toml

import (
	"github.com/BurntSushi/toml"
)

func ReadConfig() (Config, error) {
	var config Config

	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		return config, err
	}

	return config, nil
}

func ReadPrivate() (Private, error) {
	var config Private

	if _, err := toml.DecodeFile("config/private.toml", &config); err != nil {
		return config, err
	}

	return config, nil
}