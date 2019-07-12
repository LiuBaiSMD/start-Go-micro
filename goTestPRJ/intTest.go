package main

import (
	"fmt"
	"unsafe"
)

var arr = [3]int{1, 2, 3}

func main()  {
	fmt.Println(&arr[0],&arr[1])
	fmt.Println(unsafe.Sizeof(arr))
	fmt.Println(unsafe.Sizeof(arr[0]))
	i := 99999999
	fmt.Println(unsafe.Sizeof(i))
	var b *bool
	fmt.Println(unsafe.Sizeof(b))

}