# 2019.07.02
## 1.整个项目概要梳理
1.使用protobuf文件来定义API接口，所传输的数据流格式，编译proto文件，形成pb、micro代码风格文件

2.开始编写服务端->注册响应服务（注册服务名字、选择对应服务类型service、function等）->通过导入proto文件所在的包，进行继承实现API接口
总结：创建服务，初始化服务，运行服务。

3.开始编写客户端->初始化服务->调用相应接口

4.go mod:启用go mod包管理,配合goproxyio使用，解决中国特色无法下载包的问题（go mod 下载的包会放在$PATH/pkg路径下）
```
export GO111MOULD=auto
export GOPROXY=https://goproxy.io
```
在项目文件路径下执行初始化代码
```
go mod init [所依赖的自己的包路径]
```
下载拉取依赖包，不执行直接运行脚本也可
```
go build 
```

5.github 分支管理操作
```
git branch name: 创建分支
git checkout branch_name: 切换分支
git branch: 查看分支,以及当前所在分支
git merge branch_name: 合并分支到当前分支
```

6.protoc编译指令
```
protoc --proto_path=. --micro_out=. --go_out=. your.proto
```
生成对应的pb.go、micro.go文件

7.GOPATH、GOROOT区分
一个是GO自身的安装目录 GOROOT，这个目录只能放标准包，也就是”standart packages“。另一个是GOPATH，也就是工作目录，这里用来放第三方包，也就是“non-standard packages”

8.go-micro微服务构成了解


## 未完成：学习并测试带IP：端口Mirco实例