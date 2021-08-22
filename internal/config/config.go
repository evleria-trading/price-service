package config

import "time"

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	GenerationRate time.Duration `env:"GENERATION_RATE" envDefault:"1ms"`
}
