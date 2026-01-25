package ports

import (
	"context"

	"github.com/araujo-angel/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	CalculateDelivery(ctx context.Context, request domain.ShippingRequest) (domain.ShippingResponse, error)
}
