package main

import (
	"log"

	"github.com/araujo-angel/microservices/order/config"
	"github.com/araujo-angel/microservices/order/internal/adapters/db"
	"github.com/araujo-angel/microservices/order/internal/adapters/grpc"
	"github.com/araujo-angel/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Faild to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(config.GetApplicationPort(), application)
	grpcAdapter.Run()
}
