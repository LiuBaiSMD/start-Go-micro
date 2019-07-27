package main

import (
	"fmt"
)
func test_slice(sl []string){
	fmt.Printf("%v, %p, %d %d\n",sl, sl, len(sl), cap(sl))
	sl[0] = "aa"
	sl = append(sl, "d")
	//sl = append(sl, "d")
	fmt.Printf("%v, %p, %d %d\n",sl, sl, len(sl), cap(sl))
}

func lencap(a []int) {
	fmt.Printf("------->%v %d %d\n\n", a, len(a), cap(a))
}

func main() {

	sl := []string{
		"a",
		"b",
		"c",
		"e",
	}

	fmt.Printf("%v, %p, %d %d \n",sl, sl, len(sl), cap(sl))
	sl = append(sl[:1], sl[2:]...)
	fmt.Printf("%v, %p, %d %d\n",sl, sl, len(sl), cap(sl))
	test_slice(sl)
	fmt.Printf("%v, %p, %d %d\n",sl, sl, len(sl), cap(sl))
}