package main

import "fmt"

func testParams(nums ...int){
	for i := range nums{
		fmt.Println(i)
	}
}
func main(){
	testParams(1,2,3,4)
}
