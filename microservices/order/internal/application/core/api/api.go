package api

import (
	"github.com/araujo-angel/microservices/order/internal/application/core/domain"
	"github.com/araujo-angel/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
	}
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, int32, error) {
	// 1. Validar quantidade total de itens (max 50)
	var totalItems float32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}

	if totalItems > 50 {
		return domain.Order{}, 0, status.Errorf(codes.InvalidArgument, "Order with more than 50 items is not allowed. Total items: %.0f", totalItems)
	}

	// 2. Validar se todos os itens existem no estoque
	for _, item := range order.OrderItems {
		_, err := a.db.GetStockItemByProductCode(item.ProductCode)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return domain.Order{}, 0, status.Errorf(codes.NotFound, "Item with product code '%s' not found in stock.", item.ProductCode)
			}
			return domain.Order{}, 0, status.Errorf(codes.Internal, "Error checking stock for product code '%s': %v", item.ProductCode, err)
		}
	}

	// 3. Salvar pedido no banco
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, 0, err
	}

	// 4. Chamar servico de pagamento
	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		return domain.Order{}, 0, paymentErr
	}

	// 5. Se pagamento OK, chamar servico de shipping
	deliveryDays, shippingErr := a.shipping.CalculateDelivery(order)
	if shippingErr != nil {
		return domain.Order{}, 0, shippingErr
	}

	return order, deliveryDays, nil
}
