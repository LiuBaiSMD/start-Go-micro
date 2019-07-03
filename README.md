# micro学习
## 搭建micro环境
[Micro China](https://github.com/micro-in-cn)

build-start-micro.md文档中，在已经搭建好go环境的前提下，搭建go micro微服务架构
***
相关知识
***
consul：微服务发现框架，解决多服务环境下客户端相对于的配置问题

相关介绍链接：

[consul图文介绍](https://www.cnblogs.com/xiaohanlin/p/8016803.html)

[consul介绍、安装、使用教程](https://blog.csdn.net/liuzhuchen/article/details/81913562)

[consul与其他代理比较](https://www.jianshu.com/p/e0986abbfe48)

[客户端微服务发现](https://microservices.io/patterns/client-side-discovery.html)

[服务端微服务发现](https://microservices.io/patterns/server-side-discovery.html)

protobuff：是一个灵活的、高效的用于序列化数据的协议

protoc：Protobuf（Protocol Buffers - Google's data interchange format）编译器：

protoc-gen-go：goprotobuf 提供的 Protobuf 插件：在$GOPATH目录下执行go get github.com/micro/protobuf/{proto,protoc-gen-go}，该命令会在bin目录下生成protoc-gen-go(.exe)工具，protoc编译器利用protoc-gen-go插件将.proto文件转换为Golang源文件

protoc-gen-micro（Protobuf code generation for micro）：在$GOPATH目录下执行go get github.com/micro/protoc-gen-micro，该命令会在bin目录下生成protoc-gen-micro(.exe)，protoc编译器利用protoc-gen-micro插件将.proto文件转换为micro代码风格文件

[Golang下使用protobuf](https://www.jianshu.com/p/f4051569fd32)

[protocol buffers生成go代码原理](https://www.cnblogs.com/charlieroro/p/9043337.html)

# 如何创建实例
1.通过[环境安装教程](https://github.com/LiuBaiSMD/start-Go-micro/blob/master/build-start-micro.md)安装好环境

2.下载goPRJ到本地GO项目存放路径下(如使用gomod管理包，请将本项目放置在$GOPATH路径外)

3.设置goproxyio代理，启动go mod包管理，拉取依赖包
```
export GOPROXY=https://goproxy.io
export GO111MODULE=auto
go mod init goPRJ(此名字与本地项目文件夹同名即可，即proto所在的文件夹，默认为goPRJ)
go mod tidy (拉取包依赖)
```

4.并使用protoc编译指令将proto文件编译成pb.go、micro.go代码风格文件，或使用编译脚本proto_gen_recurse.sh直接编译所有proto文件
```
protoc --proto_path=. --go_out=.  --micro_out=. hello_world.proto（你的proto文件名字）
```

5.启动consul
```
consul agent -dev
```
查看进程是否启动
```
micro --registry=consul list services
```
可进入[consul WEB管理器](http://localhost:8500)查看

6.启动service端
goPRJ文件夹结构
```
├── goPRJ
    ├── Clients
    │   └── helloclient.go
    └── services
        └── hello.go
```
运行goPRJ下的hello.go
```
go run hello.go
```
***保持service后台运行***

7.启动Client进行测试
运行goPRJ下的helloclient.go
```
go run helloclient.go
```
在运行后会在界面返回对应结果
