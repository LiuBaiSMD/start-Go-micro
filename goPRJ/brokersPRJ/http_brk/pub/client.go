package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/http"
	"github.com/micro/go-micro/registry/mdns"
	"time"
)

func main(){
	msg := &broker.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}
	rgst := mdns.NewRegistry()
	subBroker := http.NewBroker(broker.Registry(rgst))
	//subBroker.
	if err := subBroker.Init(); err != nil {
		fmt.Println("Unexpected init error: %v", err)
	}

	if err := subBroker.Connect(); err != nil {
		fmt.Println("Unexpected connect error: %v", err)
	}
	for i :=0;i < 10;i++{
		subBroker.Publish("test", msg)
		time.Sleep(time.Second * 1)
		fmt.Println("i send a msg")
	}
	time.Sleep(time.Second * 100)
}

