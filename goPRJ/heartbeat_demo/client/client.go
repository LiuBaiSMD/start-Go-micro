
package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/util/log"
	"heartbeat_demo/proto"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	CLIENTID = 10
	USERID   = 10001
)

var (
	clientRes heartbeat.Request
	wsHost          = "127.0.0.1:8080"
	wsPath          = "/heartbeat"
	msgSeqId uint64 = 0
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

func (this *Client) SendMessage() error {

	// 增加一个信号监控,检测各种退出的情况,方便通知服务器断开连接
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://"+this.Host+this.Path, http.Header{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer conn.Close() //关闭连接

	done := make(chan struct{})
	// 另外其一个goroutine处理接收消息
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatalf("read:", err)
				return
			}
			if err := proto.Unmarshal(message, &clientRes); err != nil {
				log.Logf("proto unmarshal: %s", err)
			}
			log.Logf("recv: %s", clientRes)
		}
	}()
	//进行发送输入功能
	reader:= make(chan string )
	data := ""
	go func(){
		for{
			log.Logf("please input: 	")
			fmt.Scanf("%s",&data)
			reader <- data
			log.Logf("your input : %v",data)
		}
	}()
	go func(){
		d := ""
		//ticker := time.NewTicker(time.Second*5)
		//reader := make(chan string)
		for{
			select {
			//case <-ticker.C:
			//	err1 :=conn.WriteMessage(websocket.BinaryMessage, MsgAssembler())
			//	if err1 != nil {
			//		log.Printf("write close:", err1)
			//	} else {
			//		time.Sleep(time.Second * 5)
			//	}
			case d =<- reader:
				log.Logf("----->send your input")
				err1 :=conn.WriteMessage(websocket.BinaryMessage, MsgAssemblerReader(d))
				if err1 != nil {
					log.Logf("write close:", err1)
				} else {
					log.Logf("send input over!")
				}
			}

		}
	}()
	ticker := time.NewTicker(time.Second*5)
	defer ticker.Stop()
	d := ""
	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			err := conn.WriteMessage(websocket.BinaryMessage, MsgAssembler())
			if err != nil {
				log.Fatalf("write:", err)
				return nil
			}
		case d=<-reader:
			log.Logf("----->send your input")
			err1 :=conn.WriteMessage(websocket.BinaryMessage, MsgAssemblerReader(d))
			if err1 != nil {
				log.Logf("write close:", err1)
			} else {
				log.Logf("send input over!")
			}
		case <-interrupt:
			log.Fatalf("interrupt")

			// 发送 CloseMessage 类型的消息来通知服务器关闭连接，不然会报错CloseAbnormalClosure 1006错误
			// 等待服务器关闭连接，如果超时自动关闭.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatalf("write close:", err)
				return nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}


func main() {
	//// 一个简易socket连接的客户端服务
	//service := micro.NewService(
	//	micro.Name("go.micro.web.wsClient"),
	//)
	//service.Init()
	//// 这里就开始发了别影响服务启动
	//go msgHandler()
	//if err := service.Run(); err != nil {
	//	log.Logf("Run: errr %v", err)
	//}
	msgHandler()
}

// 组装pb的接口
func MsgAssembler() []byte {
	msgSeqId += 1
	retPb := &heartbeat.Request{
		ClientId: CLIENTID,
		UserId:   USERID,
		MsgId:    msgSeqId,
		Data:     "handshake:",
	}
	byteData, err := proto.Marshal(retPb)
	if err != nil {
		log.Fatal("pb marshaling error: ", err)
	}
	return byteData
}
func msgHandler() {
	clientWrapper := NewWebsocketClient(wsHost, wsPath)
	if err := clientWrapper.SendMessage(); err != nil {
		log.Logf("SendMessage: errr%v", err)
	}
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