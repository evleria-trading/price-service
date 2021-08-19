package main

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/price-service/internal/chanPool"
	"github.com/evleria/price-service/internal/config"
	"github.com/evleria/price-service/internal/generator"
	"github.com/evleria/price-service/internal/handler"
	"github.com/evleria/price-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := new(config.Config)
	check(env.Parse(cfg))

	setupLogger(cfg.Environment)

	pool := chanPool.NewChanPool()
	priceGenerator := generator.NewPricesGenerator(pool, cfg.GenerationRate)

	go func() {
		check(priceGenerator.GeneratePrices(context.Background()))
	}()

	startGrpcServer(handler.NewPriceService(pool), ":6000")
}

func startGrpcServer(service pb.PriceServiceServer, port string) {
	listener, err := net.Listen("tcp", port)
	check(err)

	s := grpc.NewServer()
	pb.RegisterPriceServiceServer(s, service)
	reflection.Register(s)

	log.Info("gRPC server started on ", port)
	check(s.Serve(listener))
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
