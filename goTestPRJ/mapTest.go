package main

import "fmt"

func main()  {
	var m  map[string]string
	m = make(map[string	]string)
	m["name"] = "wuxun"
	m["age"] = "15"

	fmt.Println(m["age"])

}
