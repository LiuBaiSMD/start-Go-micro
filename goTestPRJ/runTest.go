package main

import "fmt"

var a int32 = 1

func test1() {
	if a != 1 {
		fmt.Println(a)
		return
	}
	if a < 0 || a != 1 {
		fmt.Println(a)
	}
	fmt.Println("end")

}
func main() {
	test1()
}