package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

var C Config

func Configure() {
	err := envconfig.Process("", &C)
	if err != nil {
		slog.Error("Cannot read env",
			"err", err)
	}
}
