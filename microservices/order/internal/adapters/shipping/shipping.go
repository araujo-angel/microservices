package shipping

import (
	"context"
	"log"
	"time"

	"github.com/araujo-angel/microservices-proto/golang/shipping"
	"github.com/araujo-angel/microservices/order/internal/application/core/domain"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(
			retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			retry.WithMax(5),
			retry.WithBackoff(retry.BackoffLinear(time.Second)),
		)))

	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) CalculateDelivery(order domain.Order) (int32, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var items []*shipping.ShippingItem
	for _, item := range order.OrderItems {
		items = append(items, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    int32(item.Quantity),
		})
	}

	response, err := a.shipping.Calculate(ctx, &shipping.CalculateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})

	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout ao chamar servico de shipping para OrderID %d: deadline de 15 segundos excedido", order.ID)
			}
		}
		return 0, err
	}

	return response.DeliveryDays, nil
}
