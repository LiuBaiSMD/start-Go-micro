package main

import "fmt"

//type funcs interface {
//	func func4(name, age string)
//}
type Test struct {
	func1 func(name , sex string)
	func2 func(age int)
	func3 string
}
type Test1 struct {
	func1 func(name , sex string)
	func2 func(age int)
	func3 string
}

func func1(name , sex string){
	fmt.Println(name,sex)
}

func (T *Test) func4(name, age string){
	fmt.Println(name,age)
}

func (T *Test1) func4(name, age string){
	fmt.Println(name,age,"1")
}

func main() {
	var T Test
	T.func3 = "two funcs"
	T.func1 = func(name, sex string){
		fmt.Println(name, sex)
	}
	T.func1("wuxun", "male")
	fmt.Println(T.func3)
	var T1 Test1
	T1.func4("wuxun", "15")
	T.func4("suxun", "15")
}