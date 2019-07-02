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

protobuff：是一个灵活的、高效的用于序列化数据的协议

[Golang下使用protobuf](https://www.jianshu.com/p/f4051569fd32)

[protocol buffers生成go代码原理](https://www.cnblogs.com/charlieroro/p/9043337.html)

# 如何创建实例
1.通过[环境安装教程](https://github.com/LiuBaiSMD/start-Go-micro/blob/master/build-start-micro.md)安装好环境

2.下载goPRJ到本地GO项目存放路径下

3.下载proto路径到本地GOPATH路径下，并使用protoc编译成pb.go文件
```
protoc --go_out=plugins=micro:. hello_world.proto
```

4.启动consul
```
consul agent -dev
```
查看进程是否启动
```
micro --registry=consul list services
```
可进入[consul WEB管理器](http://localhost:8500)查看

5.启动service端
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
go run goPRJ/services/helloclien.go
```
***保持service后台运行***

6.启动Client进行测试
运行goPRJ下的helloclient.go
```
go run goPRJ/Clients/helloclient.go
```
在运行后会在界面返回结果
