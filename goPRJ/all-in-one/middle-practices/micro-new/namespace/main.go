package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"namespace/handler"
	"namespace/subscriber"
	namespace "namespace/proto/namespace"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.srv.namespace"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	namespace.RegisterNamespaceHandler(service.Server(), new(handler.Namespace))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.srv.namespace", service.Server(), new(subscriber.Namespace))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.srv.namespace", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
 