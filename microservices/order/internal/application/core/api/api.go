package api

import (
	"github.com/araujo-angel/microservices/order/internal/application/core/domain"
	"github.com/araujo-angel/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	var totalItems float32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}

	if totalItems > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Order with more than 50 items is not allowed. Total items: %.0f", totalItems)
	}

	err := a.db.Save(&order)
	if err != nil {

		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}
	return order, nil
}
