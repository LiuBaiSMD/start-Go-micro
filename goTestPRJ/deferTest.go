package main

import "fmt"

func foo(){
	fmt.Println("foo")
}

func foo1(i int) int {
	i = 100
	i = 200
	fmt.Println("foo1")
	return i
}

func foo2(i int) int {
	i = 100
	fmt.Println("foo2")
	defer foo()
	fmt.Println("foo3")
	i = 200
	return i
}

func main(){
	foo2(3)
}