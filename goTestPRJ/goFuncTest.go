package main

import (
	"fmt"
	"time"
)

func main()  {
i := 1
j := 2
go func(i , j int) {
time.Sleep(1000*time.Millisecond)
fmt.Println("i =", i, j)
} (i,j)
time.Sleep(100*time.Millisecond)

i++
time.Sleep(1000*time.Millisecond)
}

//func main()  {
//	i := 1
//
//	go func() {
//		time.Sleep(100*time.Millisecond)
//		fmt.Println("i =", i)
//	} ()
//
//	i++
//	time.Sleep(1000*time.Millisecond)
//}