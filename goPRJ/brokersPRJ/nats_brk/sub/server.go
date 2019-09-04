package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry/mdns"
	"time"
)

func main(){
	rgst := mdns.NewRegistry()
	subBroker := nats.NewBroker(broker.Registry(rgst))
	//subBroker.
	if err := subBroker.Init(); err != nil {
		fmt.Println("Unexpected init error: %v", err)
	}

	if err := subBroker.Connect(); err != nil {
		fmt.Println("Unexpected connect error: %v", err)
	}
	subBroker.Subscribe("test", pubMSG)
	time.Sleep(time.Second * 100)
}

func pubMSG(p broker.Event) error{
	fmt.Println("msg:	", p)
	return nil
}