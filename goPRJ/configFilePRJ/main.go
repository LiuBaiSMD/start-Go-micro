package main

import (
	"fmt"
	"reflect"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

func main() {
	// 加载配置文件
	if err := config.Load(file.NewSource(
		file.WithPath("/Users/tugame/Projects/goPRJ/configFilePRJ/config/config.yml"),
		//file.WithPath("/Users/tugame/Projects/goPRJ/configFilePRJ/config/config.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	// 根据实际情况，定义合适的结构
	// go-config通过scan方法将配置转成JSON，再传入指定类型的field中
	type Host struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Port    int    `json:"port"`
		Test	string `json:"Test"`
		Test1	string `json:"Test1"`
	}

	var host Host
	//if err := config.Get("hosts", "cache").Scan(&host); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	w, err := config.Watch("hosts", "cache")
	if err != nil {
		// do something
	}
	v, err := w.Next()
	if err != nil {
		// do something
	}
	v.Scan(&host)
	c := config.Map()
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(host))
	fmt.Println(host.Name, host.Address, host.Port, host.Test)
	}
