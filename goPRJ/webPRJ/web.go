package main

import (
	"fmt"
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
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// static files
	fmt.Println("html1")
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("1.html"))))
	// websocket interface
	service.HandleFunc("/websocket/hi", hi)
	fmt.Println("html2")

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
	fmt.Println("html3")

}

func hi(w http.ResponseWriter, r *http.Request) {

	log.Println("i get your request!")
	fmt.Println(r.Header)
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}
	fmt.Println("html4")
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