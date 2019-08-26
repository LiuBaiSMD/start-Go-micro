package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"net/http"
	"os"
	"time"
	"html/template"
	"github.com/go-redis/redis"
	"strconv"
	//"github.com/micro/go-web"
	"userwebPRJ/handler"
)

var rdsClient *redis.Client

var upGrader = websocket.Upgrader{
CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

	// 初始化配置
	//base.Init()

	//dockerMode = os.Getenv("RUN_DOCKER_MODE")
	//连接Redis
	redisUrl := "localhost:6379"
	connType := "tcp"
	rdsClient = connRedis(connType, redisUrl)
	fmt.Print(rdsClient)

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
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html/websocket"))))
	service.Handle("/login/", http.StripPrefix("/login/", http.FileServer(http.Dir("html/login"))))

	service.HandleFunc("/userlogin/", UserLogin)
	service.HandleFunc("/userregister/", UserRegister)
	service.HandleFunc("/websocket/hi", hi)
	service.HandleFunc("/setRedisHash", setRedisHashReq)

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

func setRedisHashReq(w http.ResponseWriter, r *http.Request){
	// Redis存储请求接口
	r.ParseForm()
	key := r.Form.Get("name")
	value := r.Form.Get("age")
	last := r.Form.Get("last")
	if(key == "" || value == ""){
		log.Logf("参数错误")
		return
	}
	expTime,err := strconv.Atoi(last)
	if err != nil{
		log.Logf("参数错误：	", last)
	} else{
		setRedisHash(key, value, rdsClient,expTime)
	}
}

func setRedisHash(key , value string, rdsConn *redis.Client, expire int){
	log.Logf("插入数据：%v  :%v", key, value)
	err := rdsConn.Set(key, value, time.Duration(time.Second * 15)).Err()
	if err != nil {
		panic(err)
	}
}

func connRedis(connType, redisUrl string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}

func UserLogin(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")
	log.Log("method:", r.Method) //获取请求的方法
	fmt.Fprintf(w, "密码错误，请重试!")
	return
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/userlogin.html")
		t.Execute(w, nil)

	} else {
		t := template.New("test")
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		name := r.Form["name"]
		password := r.Form["password"]
		if name[0] == "" || password[0] == ""{
			fmt.Fprintf(w, "请填写账号、密码！")
			return
		}
		//log.Log("name:", name)
		//log.Log("password:", password)
		ifSuccess := 0
		if ifSuccess == 0 {
			url := "/websocket/"
			http.Redirect(w,r, url, http.StatusFound)
		}else if ifSuccess == 1{
			fmt.Fprintf(w, "无此用户，请进行注册!")
			return
		}else if ifSuccess == 2{
			fmt.Fprintf(w, "密码错误，请重试!")
			return
		}
		t.Execute(w, nil)

	}
}

func UserRegister(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/register.html")
		t.Execute(w, nil)

	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		name := r.Form["name"]
		password := r.Form["password"]
		password2 := r.Form["password2"]
		fmt.Println("name:", name, password, password2)
		//t := template.New("test")
		//进行参数检查
		if name[0] == ""{
			fmt.Fprintf(w, "请输入用户名！")
			return
		}
		if password[0] == "" || password2[0] == ""{
			fmt.Fprintf(w, "密码不能为空")
			return
		}else if password[0] != password2[0]{
			fmt.Fprintf(w, "两次密码不一致！请确认后重试！")
			return
		}
		//进行数据库检查
		ifOK := 0
		if ifOK == 1{
			fmt.Fprintf(w, "密码过于短小，请重设密码!（8个字符起！）")
			return
		}else if ifOK == 2{
			fmt.Fprintf(w, "此账户已被注册！")
			return
		}else if ifOK == 0{
			url := "/websocket/"
			http.Redirect(w,r, url, http.StatusFound)
		}

	}
}

func UserLoginVue(w http.ResponseWriter, r *http.Request){
	log.Log("method:", r.Method) //获取请求的方法
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

//func hi(w http.ResponseWriter, r *http.Request) {
//	log.Logf("hi")
//
//	c, err := upGrader.Upgrade(w, r, nil)
//	log.Logf("recv: %s", c)
//	if err != nil {
//		log.Logf("upgrade: %s", err)
//		return
//	}
//	defer c.Close()
//	for {
//		mt, message, err := c.ReadMessage()
//		if err != nil {
//			log.Logf("read:", err)
//			break
//		}
//
//		log.Logf("recv: %s", message)
//		err = c.WriteMessage(mt,  message)
//		if err != nil {
//			log.Logf("write:", err)
//			break
//		}
//	}
//}