package main

import "fmt"

func testParams(nums ...int){
	testParams1(nums...)
	for i := range nums{
		fmt.Println(i)
	}
}

func testParams1(nums ...int){
	fmt.Println(nums)
}

func main(){
	testParams(1,2,3,4)
}
