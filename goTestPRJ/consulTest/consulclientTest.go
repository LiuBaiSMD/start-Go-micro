package main

import (
"context"
"fmt"

micro "github.com/micro/go-micro"
proto "consulTest/APIPT"

"github.com/micro/go-micro/registry"
"github.com/micro/go-micro/registry/consul"
)


func main() {
	reg := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	fmt.Println("start")
	service := micro.NewService(micro.Registry(reg),micro.Name("go.micro.api.client"))
	service.Init()
	greeter := proto.NewExampleService("go.micro.api.example", service.Client())

	//rsp, err := greeter.Call(context.TODO(), &proto.Request{Method:"GET" , Path:"/example/call",Url:"/example/call?name=john"})
	rsp, err := greeter.Call(context.TODO(), &proto.Request{Body: "John, how are you?"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
	fmt.Println("over")
}
