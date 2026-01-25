package api

import (
	"context"

	"github.com/araujo-angel/microservices/shipping/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a Application) CalculateDelivery(ctx context.Context, request domain.ShippingRequest) (domain.ShippingResponse, error) {
	if len(request.Items) == 0 {
		return domain.ShippingResponse{}, status.Errorf(codes.InvalidArgument, "Items list cannot be empty.")
	}

	deliveryDays := domain.CalculateDeliveryDays(request.Items)

	return domain.ShippingResponse{
		DeliveryDays: deliveryDays,
	}, nil
}
