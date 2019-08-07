package main

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"path/filepath"
)

func main(){
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	e := json.NewEncoder()
	fmt.Println(appPath)
	fileSource := file.NewSource(
		file.WithPath(appPath+"/conf/micro.yml"),
		source.WithEncoder(e),
	)
	conf := config.NewConfig()
	if err := conf.Load(fileSource); err != nil {
		log.Logf("load config errr!", err)
	}
	//strMap := make(map[string]string)
	//confulConf := conf.Get("consul").StringMap(strMap)



	//开始测试map[string]interface{} 深度拷贝
	conSIMap := conf.Map()
	fmt.Printf("\nconSIMap-------->Type: %T \n", conSIMap)
	test := conSIMap["change"]
	fmt.Println("test=======:", test)
	newCp := deepCopySI(conSIMap)
	conSIMap["change"] = "NO"
	fmt.Println("=========>>>>>>old", conSIMap)
	fmt.Printf("newCP-------->     %v ", newCp)
	fmt.Printf("\nnewCP-------->Type: %T \n", newCp)

	fmt.Println(newCp.(map[string]interface {}))
	a := newCp.(map[string]interface {})["change"]
	newCp.(map[string]interface {})["change"] = "NO and YES"
	fmt.Printf("newCP-------->     %v ", newCp)

	fmt.Println("a=======:", a)
	for k, v := range newCp.(map[string]interface {}) {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float", int64(vv))
			case int:
				fmt.Println(k, "is int", vv)
				//newCp["test"] =
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			case nil:
				fmt.Println(k, "is nil", "null")
			case map[string]interface{}:
				fmt.Println(k, "is an map:")
				//print_json(vv)
			default:
				fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))

		}
	}
	for k, v := range conSIMap {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
			//newCp["test"] =
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			//print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))

		}
	}

	//fmt.Printf("newCP-------->%v \nType: %T", newCp, newCp)
	//for k,v := range newCp["consul"]{
	//	fmt.Println(k,v)
	//}

		}



func deepCopySI(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = deepCopySI(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = deepCopySI(v)
		}

		return newSlice
	}

	return value
}