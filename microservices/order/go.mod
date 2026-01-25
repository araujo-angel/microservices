module github.com/araujo-angel/microservices/order

go 1.25.1

require (
	github.com/araujo-angel/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	github.com/araujo-angel/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
	github.com/araujo-angel/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.3
	google.golang.org/grpc v1.78.0
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/araujo-angel/microservices-proto/golang/order => ../../../microservices-proto/microservices-proto/golang/order

replace github.com/araujo-angel/microservices-proto/golang/payment => ../../../microservices-proto/microservices-proto/golang/payment

replace github.com/araujo-angel/microservices-proto/golang/shipping => ../../../microservices-proto/microservices-proto/golang/shipping
