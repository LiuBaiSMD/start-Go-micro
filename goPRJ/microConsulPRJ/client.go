package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	proto "microConsulPRJ/proto"

)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(micro.Registry(reg), micro.Name("greeter.client"))
	service.Init()
	rsp, err := proto.NewGreeterService("helloworld", service.Client()).Hello(context.TODO(), &proto.HelloRequest{Name: "benben_2015"})
	if err != nil {
		fmt.Println("failed to new greeter service: ", err)
	}
	fmt.Println(rsp.Greeting)
}
