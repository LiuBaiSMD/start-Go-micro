package main

import "fmt"

type Stu struct {
	Name  string `json:"name"`
	Age   int	`json:"age"`
	HIgh  bool `json:"High"`
	sex   string `json:"sex"`
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}


func main()  {
	a:=10
	b:=0
	c:= a/2+b
	//if err != nil{
	//	fmt.Println("出现错误")
	//}
	fmt.Println("没有出现错误", c)
	s :=Stu{}
	fmt.Println("S:", s)
}
