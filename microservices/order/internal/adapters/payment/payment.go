package payment

import (
	"context"
	"log"
	"time"

	"github.com/araujo-angel/microservices-proto/golang/payment"
	"github.com/araujo-angel/microservices/order/internal/application/core/domain"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(
			retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			retry.WithMax(5),
			retry.WithBackoff(retry.BackoffLinear(time.Second)),
		)))

	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order domain.Order) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err := a.payment.Create(ctx, &payment.
		CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout ao chamar serviço de pagamento para OrderID %d: deadline de 15 segundos excedido", order.ID)
			}
		}
	}
	return err
}
