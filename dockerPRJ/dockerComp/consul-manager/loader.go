package main

import (
	"bytes"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"sync"
)

var (
	m                      sync.RWMutex
	inited                 bool
	err                    error
	consulAddr             consulConfig
	consulConfigCenterAddr string
)

// consulConfig 配置结构
type consulConfig struct {
	Enabled    bool   `json:"enabled"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	KVLocation string `json:"kv_location"`
}

// Init 初始化配置
func Init() {
	m.Lock()
	//进行配置推送检测，是否已经推送过配置
	defer m.Unlock()
	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}

	// 加载yml默认配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	e := json.NewEncoder()
	fmt.Println(appPath)
	fileSource := file.NewSource(
		file.WithPath(appPath+"/conf/micro.yml"),
		source.WithEncoder(e),
	)
	conf := config.NewConfig()
	// 加载micro.yml文件
	if err = conf.Load(fileSource); err != nil {
		panic(err)
	}
	fmt.Println(string(conf.Bytes()))

	// 读取连接的配置中心
	configMap := conf.Map()
	fmt.Println(configMap)
	//scan将配置读入到放入的变量consulAddr之中
	if err := conf.Get("consul").Scan(&consulAddr); err != nil {
		panic(err)
	}
	// 拼接配置的地址和 KVcenter 存储路径,对本地以及docker环境进行判断
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode != "on"{
		fmt.Println("本地模式2")
		consulConfigCenterAddr = consulAddr.Host + ":" + strconv.Itoa(consulAddr.Port)
	}else{
		fmt.Println("docker模式")
		var consulService string
		if err := conf.Get("consul","docker_host").Scan(&consulService); err != nil {
			panic(err)
		}
		consulConfigCenterAddr = consulService
	}
	//进行测试
	//consulConfigCenterAddr = consulAddr.Host + ":" + strconv.Itoa(consulAddr.Port)

	url := fmt.Sprintf("http://%s/v1/kv/%s", consulConfigCenterAddr, consulAddr.KVLocation)
	fmt.Println("url:", url)
	_, err, _ := PutJson(url, string(conf.Bytes()))
	if err != nil {
		log.Fatalf("http 发送模块异常，%s", err)
		panic(err)
	}
	// 侦听文件变动
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatalf("[Init] 开始侦听应用配置文件变动 异常，%s", err)
		panic(err)
	}

	fmt.Println(consulConfigCenterAddr)
	oldStrMap := make(map[string]string)
	oldStrMap = conf.Get().StringMap(oldStrMap)
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] 侦听应用配置文件变动 异常， %s", err)
				return
			}
			if err = conf.Load(fileSource); err != nil {
				panic(err)
			}
			log.Logf("[loadAndWatchConfigFile] 文件变动，%s", string(v.Bytes()))

			////本部分代码还有部分问题 1.对于底层修改、增删的部分只会认为是change
			strMap := make(map[string]string)
			newMapConf := v.StringMap(strMap)
			findConfDif(oldStrMap, newMapConf)

			_, err, _ = PutJson(url, string(v.Bytes()))
			if err != nil {
				log.Fatalf("http 发送模块异常，%s", err)
				panic(err)
			}
			fmt.Println("配置重新上传完毕！")
			oldStrMap = deepCopy(newMapConf)
		}
	}()
	// 标记已经初始化
	inited = true
	return
}
func PutJson(url, body string) (ret string, err error, resp *http.Response) {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("PUT", url, buf)
	if err != nil {
		panic(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Log(err.Error())
		return "", err, resp
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err, resp
	}

	return string(data), nil, resp
}

func firstFind(oldConf map[string]interface{}, newConf map[string]string)(addConf map[string]interface{}, subConf map[string]interface{}, changeConf map[string]interface{}) {
	//fmt.Println("------------->oldConf", oldConf)
	//fmt.Println("------------->newConf", newConf)
	//先遍历一遍查看减少的配置
	addConf = make(map[string]interface{})
	subConf = make(map[string]interface{})
	changeConf = make(map[string]interface{})
	for key, value := range oldConf {
		//fmt.Println(key, ":", value)
		if newData, ok := newConf[key]; ok{
			if newData != value{
				changeConf[string(key)] = value
			}
		}else{
			subConf[string(key)] = value
		}
	}
	for key, value := range newConf {
		//fmt.Println(key, ":", value)
		if _, ok := oldConf[key]; !ok{
			addConf[string(key)] = string(value)
		}
	}
	fmt.Println("add---------->", addConf)
	fmt.Println("sub---------->", subConf)
	fmt.Println("change------->", changeConf)
	return addConf, subConf, changeConf
}

func findConfDif(oldConf map[string]string, newConf map[string]string)(addConf map[string]string, subConf map[string]string, changeConf map[string]string) {
	//遍历旧配置一遍查看减少的配置,和改变的配置
	addConf = make(map[string]string)
	subConf = make(map[string]string)
	changeConf = make(map[string]string)
	for key, value := range oldConf {
		if newData, ok := newConf[key]; ok{
			if newData != value{
				//在旧配置中存在却不相等的配置  changeConf
				changeConf[string(key)] = string(value)
			}
		}else{
			//旧配置中不存在的配置  subConf
			subConf[string(key)] = string(value)
		}
	}
	//遍历新配置  查看增加的配置
	for key, value := range newConf {
		//fmt.Println(key, ":", value)
		if _, ok := oldConf[key]; !ok{
			addConf[string(key)] = string(value)
		}
	}
	fmt.Println("add---------->", addConf)
	fmt.Println("sub---------->", subConf)
	fmt.Println("change------->", changeConf)
	return addConf, subConf, changeConf
}

func deepCopy(oldMap map[string]string)(newMap map[string]string ){
	//map[string]string只使用一层拷贝即可
	newMap = make(map[string]string)
	for key, value := range oldMap {
		newMap[key] = value
	}
	return newMap
}

func print_json(m map[string]interface{}) {
	var answer map[string] string
	answer = make(map[string]string)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			answer[string(k)] = "test"
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}