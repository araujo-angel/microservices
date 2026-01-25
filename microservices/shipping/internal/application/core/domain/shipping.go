package domain

type ShippingItem struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type ShippingRequest struct {
	OrderID int64          `json:"order_id"`
	Items   []ShippingItem `json:"items"`
}

type ShippingResponse struct {
	DeliveryDays int32 `json:"delivery_days"`
}

func NewShippingRequest(orderID int64, items []ShippingItem) ShippingRequest {
	return ShippingRequest{
		OrderID: orderID,
		Items:   items,
	}
}

// CalculateDeliveryDays calcula o prazo de entrega baseado na quantidade total de unidades.
// Prazo minimo: 1 dia. A cada 5 unidades, adiciona 1 dia.
func CalculateDeliveryDays(items []ShippingItem) int32 {
	var totalUnits int32 = 0

	for _, item := range items {
		totalUnits += item.Quantity
	}

	if totalUnits <= 0 {
		return 1
	}
	deliveryDays := (totalUnits + 4) / 5

	if deliveryDays < 1 {
		deliveryDays = 1
	}

	return deliveryDays
}
