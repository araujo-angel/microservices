package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/araujo-angel/microservices-proto/golang/shipping"
	"github.com/araujo-angel/microservices/shipping/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Calculate(ctx context.Context, request *shipping.CalculateShippingRequest) (*shipping.CalculateShippingResponse, error) {
	log.Printf("Calculating shipping for order %d...", request.OrderId)

	// Converter proto items para domain items
	var items []domain.ShippingItem
	for _, item := range request.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingRequest := domain.NewShippingRequest(request.OrderId, items)
	result, err := a.api.CalculateDelivery(ctx, shippingRequest)

	code := status.Code(err)
	if code == codes.InvalidArgument {
		return nil, err
	} else if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to calculate shipping. %v", err)).Err()
	}

	return &shipping.CalculateShippingResponse{
		DeliveryDays: result.DeliveryDays,
	}, nil
}
