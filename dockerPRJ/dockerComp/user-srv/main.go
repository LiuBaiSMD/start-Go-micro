package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/LiuBaiSMD/goProPlugins/user/base"
	"github.com/LiuBaiSMD/goProPlugins/user/base/config"
	proto "github.com/LiuBaiSMD/goProPlugins/user/proto/user"
	"time"
	"user-srv/handler"
	"user-srv/model"
	"os"
)

var (
	dockerMode string
)

func main() {

	// 初始化配置、数据库等信息
	base.Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("bambooRat.micro.srv.user"),
		micro.RegisterTTL(time.Second*15),      // 指定TTL
		micro.RegisterInterval(time.Second*10), //指定重新注册的间隔
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	proto.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	//
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	log.Logf("consul config: ", consulCfg.GetDockerHost(),consulCfg.GetHost(), consulCfg.GetPort())
	if dockerMode == "on" {
		log.Logf("docker模式")
		ops.Addrs = []string{fmt.Sprintf("%s:%d", "consul4", 8500)}
	} else {
		log.Logf("本地模式")
		ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 8500)}
		//ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
	}
}
