# 2019.07.03
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

# 2019.07.04
1.启用Micro API测试Micro实例程序，调试在RPC、get、post方式下方法调用
```
go micro api --handler=api
```
如使用rpc模式启动,则不能使用get、post调用service服务
```
go micro api --handler=rpc
```
2.熟悉go micro工程结构，proto文件作用以及编写规则

3.了解RPC调用Micro方法时，调用中proto中的方法编写

4.可以启动多个service服务，会采用默认的负载均衡策略

5.service.go文件中需要在main方法进行实例化，可以在一个service.go文件中创建多个service服务，只需要继承实现proto中相应的service服务API接口即可

# 2019.07.08
1.通过教程将Micro的启动模式rpc、api、http（proxy）、web、event、meta处理方式进行实例启动（具体参考[micro-api介绍文档](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api)）

2.学习了解go mod使用

3.go代码的编译以及包的寻找路径、先后顺序

4.
相对路径    
```
import   "./model"  //当前文件同一目录的model目录，但是不建议这种方式import
```
绝对路径    
```
import   "shorturl/model"  //加载GOPATH/src/shorturl/model模块
```

5.导入包的查找路径
```
1.$GOROOT
2.$GOPATH
```

6.初始化路径
[初始化顺序](https://blog.csdn.net/newdas123/article/details/81082392)

7.不要无脑尝试，寻找方法比漫无目的投机取巧更重要

## 未完成：实例阅读、分析，总结代码结构调整

# 2019.07.09
1.Micro API各种handler启动方法实例运行
有的代码结构基本一致，只是不同的handler启动方式有所不同

2.RPC [RPC introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/rpc)
通过RPC向go-micro应用转送请求，通常只传送请求body，头信息不封装。只接收POST请求
RPC模式下API只接收POST方式的请求，并且只支持内容格式content-type为application/json或者application/protobuf。

3.API [API introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/api)
与rpc差不多，但是会把完整的http头封装向下传送，不限制请求方法

4.http [http introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/proxy)
与http差不多，但是支持websocket

5.event [event introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/event)
代理event事件服务类型的请求

6.meta [meta introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/meta)
元数据，通过在代码中的配置选择使用上述中的某一个处理器,运行API网关
可以看到，API启动时，并没有声明handler模式，故而使用的RPC模式。所以Meta API其实是在RPC模式的基础上，通过在接口层声明端点元数据而指定服务的

7.web [web introduction](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api/web)
与http差不多，但是支持websocket

8.go time包 
通过使用time.After定时阻塞，以及time.NewTicker(time.Second)进行周期性操作，实现一个定时结束的函数处理
```
package main

import (
	"time"
	"fmt"
)

func main()  {
	tchan := time.After(time.Second*3)
	fmt.Printf("tchan type=%T\n",tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=",<-tchan)
	fmt.Println("mark 2")
}
```

```
package main

import (
	"time"
	"fmt"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		fmt.Println(time.Now().String())
		//fmt.Println(1)

		i++
	}
}

func main(){
	pub()
}
```
9.go micro broker [broker实例](https://github.com/micro-in-cn/all-in-one/blob/master/basic-practices/micro-broker/basic/main.go)
go micro 的发布订阅类型，启用一个全局变量进行发布订阅消息沟通

# 2019.07.10
1.开始阅读micro broker源码
发布订阅双方商定一个共同的topic作为信号，当以一个topic信号发布消息时，监听该topic的订阅方将受到信号，并可以使用对应的处理函数方法。

2.Broker接口
```
Broker
type Broker interface {
    Options() Options
    Address() string
    Connect() error ///启动broker服务
    Disconnect() error ///关闭Broker服务
    Init(...Option) error
    Publish(string, *Message, ...PublishOption) error  ///publish topic message
    Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)  ///注册 topic message 的 subscribe
    String() string
}
```
Connct:启动一个broker监听，是否有人注册或者订阅topic
publish:发现topic的相关服务，组装message以及body

3.golang中  "..."用法
func test1(args ...string) { //可以接受任意个string参数}

4.defer 用法
1.defer后面必须是函数调用语句，不能是其他语句，否则编译器会出错
2.defer后面的函数在defer语句所在的函数执行结束的时候会被调用
3.对象锁的自动释放
4.注意0：如何让defer函数在宿主函数的执行中间执行
5.注意1：多个defer的执行顺序
如果函数里面有多条defer指令，他们的执行顺序是反序，即后定义的defer先执行
6.注意2：defer函数参数的计算时间点
defer函数的参数是在defer语句出现的位置做计算的，而不是在函数运行的时候做计算的，即所在函数结束的时候计算的。(参数是defer语句的实时位置，而局部变量是在运行时取值的)

5.tag使用
```
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```
1.如果一个域不是以大写字母开头的，那么转换成json的时候，这个域是被忽略的
2.如果没有使用json:"name"tag，那么输出的json字段名和域名是一样的
3.总结一下，json:"name"格式串是用来指导json.Marshal/Unmarshal，在进行json串和golang对象之间转换的时候映射字段名使用的。

6.方法的参数使用interface{}代表可以使用任意类型参数
```
func validateStruct(s interface{}) bool
```