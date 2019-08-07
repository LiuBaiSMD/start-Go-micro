package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
	"heartbeat_demo/proto"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	//对请求头进行检查
	//CheckOrigin: func(r *http.Request) bool { return true },
}
var (
	clientRes heartbeat.Request
	msgSeqId uint64 = 0
	USERID uint64 = 666
	CLIENTID uint64 = 678

)

func main() {
	// New web service

	service := web.NewService(
		web.Name("go.micro.web.heartbeat"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}
	// websocket 连接接口 web.name注册根据.分割路由路径，所以注册的路径要和name对应上
	service.HandleFunc("/heartbeat", conn)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

func conn(w http.ResponseWriter, r *http.Request) {
	//
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer conn.Close()
	reader := make(chan string ,1)
	data := ""
	go func(){
		for{
			log.Printf("please input: 	")
			fmt.Scanf("%s",&data)
			reader <- data
			log.Printf("your input : %v",data)
		}
	}()
	go func(){
		d := ""
		for{
			select {
			case d =<- reader:
				log.Printf("----->send your input")
				err1 :=conn.WriteMessage(websocket.BinaryMessage, MsgAssemblerReader(d))
				if err1 != nil {
					log.Printf("write close:", err1)
				} else {
					log.Printf("send input over!")
				}
			}

		}
	}()
	for {
		_, buffer, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if err := proto.Unmarshal(buffer, &clientRes); err != nil {
			log.Printf("proto unmarshal: %s", err)
		}
		log.Printf("recv userId=%d MsgId=%d Data=%s", clientRes.UserId, clientRes.MsgId, clientRes.Data)
	}
}

// 组装pb的接口
func MsgAssembler() []byte {
	msgSeqId += 1
	retPb := &heartbeat.Request{
		ClientId: CLIENTID,
		UserId:   USERID,
		MsgId:    msgSeqId,
		Data:     "server handshake:",
	}
	byteData, err := proto.Marshal(retPb)
	if err != nil {
		log.Fatal("pb marshaling error: ", err)
	}
	return byteData
}

func MsgAssemblerReader(data string) []byte {
	msgSeqId += 1
	retPb := &heartbeat.Request{
		ClientId: CLIENTID,
		UserId:   USERID,
		MsgId:    msgSeqId,
		Data:     data,
	}
	byteData, err := proto.Marshal(retPb)
	if err != nil {
		log.Fatal("pb marshaling error: ", err)
	}
	return byteData
}