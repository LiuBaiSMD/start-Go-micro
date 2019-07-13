package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	proto "microConsulPRJ/proto"
	"time"
)

type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	//fmt.Println("i get you: ", req.Name)
	rsp.Greeting = "hello " + req.Name
	return nil
}

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("helloworld"),
		micro.RegisterTTL(time.Second * 5),
		micro.RegisterInterval(time.Second *4),
		)

	service.Init()
	err := proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		fmt.Println("failed to register a handler: ", err)
	}

	if err = service.Run(); err != nil {
		fmt.Println("failed to run a service: ", err)
	}
}
