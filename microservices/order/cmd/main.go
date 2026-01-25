package main

import (
	"log"

	"github.com/araujo-angel/microservices/order/config"
	"github.com/araujo-angel/microservices/order/internal/adapters/db"
	"github.com/araujo-angel/microservices/order/internal/adapters/grpc"
	"github.com/araujo-angel/microservices/order/internal/adapters/payment"
	"github.com/araujo-angel/microservices/order/internal/adapters/shipping"
	"github.com/araujo-angel/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to connect to payment service. Error: %v", err)
	}
	shippingAdapter, err := shipping.NewAdapter(config.GetShippingServiceURL())
	if err != nil {
		log.Fatalf("Failed to connect to shipping service. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)
	grpcAdapter := grpc.NewAdapter(config.GetApplicationPort(), application)
	grpcAdapter.Run()
}
