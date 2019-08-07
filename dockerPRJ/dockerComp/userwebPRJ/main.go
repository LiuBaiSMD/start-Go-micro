package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"os"

	//"github.com/micro/go-web"
	"userwebPRJ/handler"
	"net/http"
)


var upGrader = websocket.Upgrader{
CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

	// 初始化配置
	//base.Init()

	//dockerMode = os.Getenv("RUN_DOCKER_MODE")
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("websocket"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8080"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}
	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))
	service.HandleFunc("/websocket/hi", hi)
	//service.HandleFunc("/websocket/hi", handler.hi)
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("over")
}

func registryOptions(ops *registry.Options) {
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		fmt.Println("docker模式")
		ops.Addrs = []string{"consul1"}
	}else{
		fmt.Println("本地模式")
		ops.Addrs = []string{"127.0.0.1:8500"}
	}
}
func hi(w http.ResponseWriter, r *http.Request) {
	log.Logf("hi")

	c, err := upGrader.Upgrade(w, r, nil)
	log.Logf("recv: %s", c)
	if err != nil {
		log.Logf("upgrade: %s", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Logf("read:", err)
			break
		}

		log.Logf("recv: %s", message)
		err = c.WriteMessage(mt,  message)
		if err != nil {
			log.Logf("write:", err)
			break
		}
	}
}