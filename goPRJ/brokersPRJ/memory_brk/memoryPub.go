package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/memory"

	//"github.com/micro/go-micro/broker/memory"
	"time"
)
func main() {
	msg := &broker.Message{
		Header: map[string]string{
			"id": fmt.Sprintf("%d", 1),
		},
		Body: []byte(fmt.Sprintf("%d: %s", 2, time.Now().String())),
	}
	serverSub := memory.NewBroker()
	if err2 := serverSub.Connect(); err2 != nil {
		fmt.Println("Unexpected connect error %v", err2)
	}
	//serverPub := memory.NewBroker()
	//fmt.Println(serverPub)
	_, err := serverSub.Subscribe("test", pubMSG)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
	err1 := serverSub.Publish("test", msg)
	if err1 != nil {
		fmt.Println(err1)
	}

}

func pubMSG(p broker.Event) error{
	fmt.Println("msg:	", p)
	return nil
}

