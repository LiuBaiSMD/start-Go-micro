package main

import (
    "fmt"
    proto "goPRJ/proto/greeterPT"
    "github.com/micro/go-micro"
    "golang.org/x/net/context"
)

func main() {
    function := micro.NewService(
        micro.Name("greeter"),
        micro.Version("latest"),
    )

greeter := proto.NewGreeterService("greeter", function.Client())

// request the Hello method on the Greeter handler
rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
    Name: "John",
})
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(rsp.Greeting)
}



// create the greeter client using the service name and client
