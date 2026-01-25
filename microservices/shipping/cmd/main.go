package main

import (
	"log"

	"github.com/araujo-angel/microservices/shipping/config"
	"github.com/araujo-angel/microservices/shipping/internal/adapters/grpc"
	"github.com/araujo-angel/microservices/shipping/internal/application/core/api"
)

func main() {
	log.Println("Starting shipping service...")

	application := api.NewApplication()
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
