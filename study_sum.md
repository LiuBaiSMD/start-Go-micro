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

# 2019.07.11
1.了解micro config使用读取配置
```
if err := config.Load(file.NewSource(
		file.WithPath("./config/config.yml"),
		file.WithPath("./config/config.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}
```

2.通过encoding和reader方法将不同格式的配置统一化并且合并，
Encoder负责资源配置编码、解码。后台资源可能会存在不同的格式，编码器负责处理不同的格式，默认的格式是Json。
编码器支持以下格式：
json、yaml、toml、xml、hcl
Reader
Reader负责把多个changeset集合并成一个可查询的值集。
Reader复用Encoder编码器将changeset集解码成map[string]interface{}格式，然后将它们合成一个changeset。
它通过格式来确定使用哪个解码器。合成的changeset中的Values可以转成Go类型值，而如果有值不能加载时，其中的值也可以作为回退值使用。

3.Encoder和Reader有什么不同？
Encoder编码器负责的是后台资源数据的编解码工作。而Reader则使用不同的encoder解码，解码的配置源可能有不同的格式，而这些encoder会解决这个事情并合并成单一编码格式。
如果是文件资源，则配置的解码器取决于文件的扩展名。
如果是基于consul的配置、etcd或类似的键值对资源，则可能会从前缀中带有多个键（特定规则）的加载，也就是说资源需要明白编码，才能返回单一的变更集

4.一个变量的大小跟机器的硬件有关，与语言无太大关系

# 2019.07.12
1、Micro是一个专注于简化分布式系统开发的微服务生态系统。
```
go-micro - A pluggable Go RPC framework for writing a microservice; service discovery, client/server rpc, pub/sub, etc.

go-plugins - Plugins for go-micro including etcd, kubernetes, nats, rabbitmq, grpc, etc.

micro - A microservice toolkit containing traditional entry points; API Gateway, CLI, Slack Bot, Sidecar and Web UI.
```

2.梳理Micro思维导图
[Micro思维导图](https://app.yinxiang.com/shard/s33/nl/25249978/69093c36-0977-487e-8014-65318af620b0/)
go Micro (由go-micro、micro、go-plugins、protobuf等组成)
go-micro源码集、
micro工具集、
go-plugins插件集

3.Micro学习笔记 
[Micro学习笔记](http://note.youdao.com/noteshare?id=91f423a06568a35e8e16cae7e1e1aeb2)

4.go 学习笔记
[go 学习笔记](http://note.youdao.com/noteshare?id=f7a5183efdc2c1e23be4666e2a80923c)

5.在启动服务时可以指定服务的TTL、interval等信息

6.通过micro new指令可以快速拉取一个go micro代码模板

7.registry注册服务、selector服务发现、broker发布订阅等

# 2019.07.13
1.consul启动，micro service、micro api、consul联合启动 [consul micro 实例](https://github.com/LiuBaiSMD/start-Go-micro/tree/master/goPRJ/microConsulPRJ)
```
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(
	micro.Registry(reg),
	micro.Name("helloworld"),
	micro.RegisterTTL(time.Second * 5),
	micro.RegisterInterval(time.Second *4),
	)
```

2.grpc实例[grpc官方实例代码](https://github.com/LiuBaiSMD/start-Go-micro/tree/master/goPRJ/grpcPRJ)

# 2019.07.14
1.开始尝试go的slackAPI接口使用
slack项目的接口API使用封装接口[nlopes/slack](https://github.com/nlopes/slack)

2.slack中bot的ID可以通过网页链接查看

3.channel的ID可以通过网页链接查看

4.用户token获取网站
[user token生成网页](https://api.slack.com/custom-integrations/legacy-tokens)

5.bot的token生成网页
[bot token生成网页](https://bamboorat.slack.com/services/你的机器人ID)

6.使用API给channel发送信息
[测试](https://github.com/LiuBaiSMD/start-Go-micro/blob/master/goPRJ/slackApiPRJ/sed2Channel.go)

7.使用API查询用户信息
[测试](https://github.com/LiuBaiSMD/start-Go-micro/blob/master/goPRJ/slackApiPRJ/getUserInfo.go)

8.使用API查询用户所在组信息
[测试](https://github.com/LiuBaiSMD/start-Go-micro/blob/master/goPRJ/slackApiPRJ/getGroupInfo.go)

9.slack API接口网站
[slack API统计](https://www.npmjs.com/package/slack)

# 2019.07.16
1.slack 接口api使用汇总
[slack API](http://note.youdao.com/noteshare?id=dcf96096dcd325d6fee5fff80e58da71)

# 2019.07.17
1.slack token获取
[获取code链接](https://api.slack.com/docs/oauth)
通过设置好
```
https://slack.com/oauth/authorize
client_id - 在您创建应用时发出（必填）
scope - 请求权限（见下文）（必填）
redirect_uri - 重定向到的URL（见下文）（可选）
state - 完成后传回的唯一字符串（可选）
team - 要尝试限制的工作空间的Slack团队ID（可选）
```
相应参数即可

# 2019.7.23
[学习笔记](http://note.youdao.com/noteshare?id=92279b39a29401faf68d04a8b6612852)
1.学习micro中go-plugins中etcd插件
从中学习正向代理与反向代理的区别、iso、TCP/IP的协议以及各层分布

2.输入url到页面加载流程

3.ARP、RARP区别

4.http中长连接和websocket的长连接的区别
[介绍与区别分析](https://www.cnblogs.com/Catherine001/p/8359153.html)

5.http和socket之长连接和短连接区别
[http和socket之长连接和短连接区别](https://www.jianshu.com/p/b68d2b26f5f4)

# 2019.7.24
1.channel使用

[channel介绍](https://studygolang.com/articles/9532)

2.转正答辩ppt准备

3.go的大小写变量代表private、public

# 2019.7.25
1.select用法
一个select语句用来选择哪个case中的发送或接收操作可以被立即执行。它类似于switch语句，但是它的case涉及到channel有关的I/O操作。

2.go-plugins 插件
1 ETCD

[etcd基本原理介绍以及雪花算法](https://www.jianshu.com/p/3bd041807974)、
[etcd基本介绍以及使用场景](https://blog.csdn.net/bbwangj/article/details/82584988)
2 Kubernets

[k8s核心概念](http://www.dockone.io/article/932)
3 NATS

[概要介绍](https://www.jianshu.com/p/2110d4f43d04)

3.go语言time模块使用

4.go struct的方法绑定

# 2019.7.26
1.双向心跳的实现

```
1.srver端使用普通的micro创建websocket服务端，对请求的域名与方法进行绑定即可，再在绑定的方法中可以使用多个协程或者调用不同的方法处理不同事务
2.client端使用普通的websocket连接即可，设置连接的Ip以及域名，进行连接，通过多个协程进行处理不同的事务。协程间可以通过channel通信。
```

2.select用法
一个select语句用来选择哪个case中的发送或接收操作可以被立即执行。它类似于switch语句，但是它的case涉及到channel有关的I/O操作。
如果同时多个case chan到达，所有的case在程序不中断的情况下会执行所有的case

3.time.NewTimer计时器
Go语言的定时器实质是单向通道，time.Timer结构体类型中有一个time.Time类型的单向chan
type Timer struct {
    C <-chan Time
    r runtimeTimer
}
time.After:指定一段时间后释放一个channel，不会阻塞后面的进程。
tchan := time.NewTimer(time.Second*3)	：三秒后放入一个信息
tchan1 := time.After(time.Second*3)		：三面后放入一个信息，实际调用的是一个NewTimer，返回一个time.Time类型的通道消息
tchan2 := time.NewTicker(time.Second * 3)：每隔三秒放入一个信息
ticker类型:一个循环计时器，每隔一段时间进行一次channle<-,

```
// Tick无法取消
      tick := time.Tick(1 * time.Minute)
      for _ = range tick {
            // do something
      }
 // 可通过调用ticker.Stop取消
      ticker := time.NewTicker(1 * time.Minute)
      for _ = range tick {
            // do something
      }
```

# 2019.7.27
1.golang中map、channel、slice在作为参数传递是的注意事项
map channel:
在方法外声明、构造，在方法内使用修改可以传递到方法外
而类似于map在方法外声明，方法内构造是无法传递到方法外
2.golang 中slice的操作用法
slice构造：
```
var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var slice = array[1:7]
slice2 := make([]int, 5, 10)
slice3 := []int{1, 2, 3, 4, 5, 6}
//4.基于数组切片创建数组切片
slice5 := slice3[:4]
```

var 变量名 []类型  //这样没有初始化赋值，仅仅是引用，没分配底层数组。
var 变量名 = []类型{置集合} //会分配底层数组,len、cap都是置集合大小
var 变量名 []类型 = make([]类型,len,cap) //这样会分配底层数组

```
type slice struct { 
    array unsafe.Pointer
    len   int
    cap   int
}
```
添加操作：s = append(s,type)
删除操作：s = append(s[0:2],s[3:]...)  	//通过头尾向append进行操作
插入操作：rear:=append([]int{},s[index:]...)
     ss=append(ss[0:index], insert)
ss= append(ss,rear...)

3.对于多个进程同时操作一个数据，通过使用channel存放数据，实现数据的交互，channel使用无缓冲的

4.回调函数的理解
[回调函数介绍](https://www.zhihu.com/question/19801131/answer/27459821)
1.定义
回调函数就是一个通过函数指针调用的函数。如果你把函数的指针（地址）作为参数传递给另一个函数，当这个指针被用来调用其所指向的函数时，我们就说这是回调函数。回调函数不是由该函数的实现方直接调用，而是在特定的事件或条件发生时由另外的一方调用的，用于对该事件或条件进行响应。
2.机制
定义一个回调函数
提供函数实现的一方在初始化的时候，将回调函数的函数指针注册给调用者
当特定的事件或条件发生时，调用者使用函数指针调用回调函数对事件进行处理
回调函数能够更加灵活的使用解决办法来处理，例如一个方法中有一个处理数据的功能，则通过将此功能设为一个回调函数，在不同场景是愚公不同的回调函数则可完成处理。

5.协程
```
1.把当前协程的 CPU 寄存器状态保存起来
2.将需要切换进来的协程的 CPU 寄存器状态加载的 CPU 寄存器上
3.协程而且完全在用户态进行
4.协程与线程主要区别是它将不再被内核调度，而是交给了程序自己而线程是将自己交给内核调度，所以也不难理解golang中调度器的存在。
5.协程是在语言层面实现对线程的调度，避免了内核级别的上下文消耗。
```

# 2019.7.28
1.进程、线程、协程、管程、纤程的介绍
[介绍连接](https://zhuanlan.zhihu.com/p/26757689)

2.进程、线程切换的理解
1、进程切换：①：代码的的加载，需要将代码从硬盘加入内存，或者切换页目录
			②：切换内核栈以及硬件上下文
2、线程切换：①：只需奇幻内核栈以及硬件上下文，
			②：在同一个虚拟内存中完成切换，所以不需要切换页目录
这两种切换都是通过操作系统内核完成，寄存器的内容切换；对于缓存的内容也要重新进行缓存，导致内存访问低效。

# 2019.7.29
1.使用Dockerfile创建docker镜像
[dockerfile配置的解释](https://blog.csdn.net/rznice/article/details/52211620)
[dockerfile使用实例](https://www.jianshu.com/p/10ed530766af)

2.dockerfile其中的from指需要整合的包的版本。

3.匿名函数与闭包
1.闭包直接使用参数的地址进行传递，使用的变量是外部实时的值；而通过值传递进入的为调用函数时保存的值，不是实时变化的值。
```
闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环
境（作用域）。
 闭包的价值
闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
```

# 2019.8.10
1.在海外服务器搭建websocket服务，并能够对请求进行转发

2.nginx配置转发，对于websocket服务需要使用upgrade请求头

3.nginx安装的路径在./configure中指定
```
./configure --prefix=/usr/local/webserver/nginx --with-http_stub_status_module --with-http_ssl_module --with-pcre=/usr/local/src/pcre-8.35
```

# 2019.8.12
1.对于docker-compose，如果需要通过localhost访问本机docker服务，需要设置network_mode:host,加入到宿主机器的本地网络段中，即可通过lcoalhost:port方式访问其他服务

2.websocket服务中，需要制定upgrade让socekt升级为websocket协议

3.github使用
```
查看以及git add的文件
git status
git reset HEAD file
git status 先看一下add 中的文件 
git reset HEAD 如果后面什么都不跟的话 就是上一次add 里面的全部撤销了 
git reset HEAD XXX.html 就是对某个文件进行撤销了
git reset --soft HEAD^   # 回滚conmmit但是不回滚代码
```

4.查看git修改的文件,并高亮
git diff .|egrep "^---|\+\+\+" --color=auto

5.git回滚到上一个版本（太危险，会直接回滚代码）
git reset --hard HEAD^         回退到上个版本
git reset --hard HEAD~3        回退到前3次提交之前，以此类推，回退到n次提交之前
git reset --hard commit_id     退到/进到 指定commit的sha码
