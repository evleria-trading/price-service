package config

import "time"

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	RedisPass string `env:"REDIS_PASS" envDefault:""`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`

	GenerationRate time.Duration `env:"GENERATION_RATE" envDefault:"250ms"`
}
