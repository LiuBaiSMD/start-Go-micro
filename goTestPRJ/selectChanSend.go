package main

import (
	"fmt"
	"time"
)

//func main() {
//	var c1, c2, c3 chan int
//	var i1, i2 int
//	select {
//	case i1 = <-c1:
//		fmt.Printf("received ", i1, " from c1\n")
//	case c2 <- i2:
//		fmt.Printf("sent ", i2, " to c2\n")
//	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
//		if ok {
//			fmt.Printf("received ", i3, " from c3\n")
//		} else {
//			fmt.Printf("c3 is closed\n")
//		}
//	default:
//		fmt.Printf("no communication\n")
//	}
//}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) { <-ch }(ch1)
	go func(ch chan int) { ch <- 2 }(ch2)

	time.Sleep(time.Second)

	for {
		select {
		case ch1 <- 1:
			fmt.Println("Send operation on ch1 works!")
		case <-ch2:
			fmt.Println("Receive operation on ch2 works!")
		default:
			fmt.Println("Exit now!")
			return
		}
	}
}