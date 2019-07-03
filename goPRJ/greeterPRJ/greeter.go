package main

import (
        "log"

        "github.com/micro/go-micro"
        proto "goPRJ/proto/greeterPT"
        "golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
        rsp.Greeting = "Hello " + req.Name
        return nil
}

func main() {
        function := micro.NewFunction(
                micro.Name("greeter"),
                micro.Version("latest"),
        )

        function.Init()

    function.Handle(new(Greeter))

        if err := function.Run(); err != nil {
                log.Fatal(err)
        }
}