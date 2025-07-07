package config

import (
	"context"
	"log"

	"github.com/caarlos0/env"
)

var Config Cfg

type Cfg struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func Parse(ctx context.Context) {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
}
