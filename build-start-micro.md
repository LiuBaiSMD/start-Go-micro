# 1. 安装micro和所需依赖
```
go get -u -v github.com/micro/micro
```
安装过程中因为长城的原因，会有很多依赖下载失败，例如x/net,x/text,x/crypt,grcp等，

**方法1.自行通过GitHub找出对应的库手动添加
方法2.通过go mod包管理工具，对墙掉的包进行依赖路径更换
方法3.通过挂起go proxy代理进程对需要的**
# 2. 安装consul
```
brew install consul
```
# 3. 安装Protobuf
```
brew install protobuf
go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u -v github.com/micro/protoc-gen-micro
```
# 4. 安装grpc和genproto（如使用go mod和goproxy代理则可不用此步骤）
```
mkdir $GOPATH/src/google.golang.org
cd $GOPATH/src/google.golang.org
git clone https://github.com/grpc/grpc-go.git grpc
git clone https://github.com/google/go-genproto.git genproto
```
# 5. 安装 golang的net,crypt,text库（如使用go mod和goproxy代理则可不用此步骤）
```
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
git clone https://github.com/golang/crypto.git
git clone https://github.com/golang/text.git
```
# 6. 最后安装micro
```
go install github.com/micro/micro
```
# 7.查看micro是否安装
```
micro --version
```
# 8.查看运行微服务进程
```
micro list servic
```
# 可能出现的问题
***
1.command-line-arguments
***
问题原因为该放置在GOPATH下的package放置在GOROOT下
GOROOT中放置的是go系统包，GOPATH放置的自行下载的包，对于包存放的位置需要自行调整。
***

2.启动consul后无法查看到进程
***
通常使用下面代码查看一般微服务进程
```
micro list servic
```
如果上述代码无法查看，是指使用consul进行服务发现和注册
```
micro --registry=consul list services
```

***
3.拉取包、库超时
***
在项目中启用go mod 与goproxy代理结合即可解决
在项目目录下使用go mod 进行包管理
```
export GOPROXY=https://goproxy.io
export GO111MODULE=auto
go mod init goPRJ(此名字与本地项目文件夹同名即可，即proto所在的文件夹，默认为goPRJ)
go mod tidy (拉取包依赖)
```