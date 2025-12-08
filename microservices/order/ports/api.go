package ports

import "github.com/araujo-angel/microservices/internal/aplication/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
