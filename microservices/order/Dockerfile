FROM golang:1.25 AS builder

RUN git clone https://github.com/araujo-angel/microservices-proto.git /microservices-proto

WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/order ./order
CMD ["./order"]
