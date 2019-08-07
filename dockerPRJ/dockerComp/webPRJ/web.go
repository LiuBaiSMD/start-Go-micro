package main

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
		"net/http"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// New web service
	//使用consul进行注册
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := web.NewService(
		web.Name("websocket"),
		web.Registry(reg),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// static files
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))

	// websocket interface
	service.HandleFunc("/websocket/hi", hi)
	//http.Handle("/websocket/hi",hi)
	fmt.Println(http.ListenAndServe(":8500",nil))
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

func hi(w http.ResponseWriter, r *http.Request) {
	log.Printf("hi")

	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
