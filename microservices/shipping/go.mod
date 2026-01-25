module github.com/araujo-angel/microservices/shipping

go 1.25.1

require (
	github.com/araujo-angel/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.78.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/araujo-angel/microservices-proto/golang/shipping => ../../../microservices-proto/microservices-proto/golang/shipping
