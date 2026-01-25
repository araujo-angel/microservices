package ports

import (
	"github.com/araujo-angel/microservices/order/internal/application/core/domain"
)

type ShippingPort interface {
	CalculateDelivery(order domain.Order) (int32, error)
}
