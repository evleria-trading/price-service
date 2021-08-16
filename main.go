package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/price-service/internal/config"
	"github.com/evleria/price-service/internal/generator"
	"github.com/evleria/price-service/internal/producer"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := new(config.Config)
	check(env.Parse(cfg))

	setupLogger(cfg.Environment)

	redisClient := getRedis(cfg)
	defer redisClient.Close()

	pricesProducer := producer.NewPriceProducer(redisClient)
	priceGenerator := generator.NewPricesGenerator(pricesProducer, cfg.GenerationRate)

	log.Info("Starting prices generation at rate: ", cfg.GenerationRate)
	check(priceGenerator.GeneratePrices(context.Background()))
}

func setupLogger(environment string) {
	switch environment {
	case "prod":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func getRedis(cfg *config.Config) *redis.Client {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
	}

	redisClient := redis.NewClient(opts)
	_, err := redisClient.Ping(context.Background()).Result()
	check(err)

	return redisClient
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
