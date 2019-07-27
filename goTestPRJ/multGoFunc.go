package main

import (
	"fmt"
	"time"
	//"math/rand"
)

func func1(){
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	chan4 := make(chan int)
	chan5 := make(chan int)
	time.Sleep(time.Second)
	fmt.Println("start")
	go func(){
		//for{
		//	rand.Seed(time.Now().UnixNano())
			//step2：获取随机数
			//num := rand.Intn(4) + 1
			chan1<-1
			chan2<-2
			chan3<-3
			chan4<-4
			chan5<-5
			//if num == 1{
			//	chan1<-num
			//}
			//if num == 2{
			//	chan2<-num
			//}
			//if num == 3{
			//	chan3<-num
			//}
			//if num == 4{
			//	chan4<-num
			//}
			//if num == 5{
			//	chan5<-num
			//}
			fmt.Println("chans over")
		//}
	}()
	for{
		select {
		case <-chan1:
			fmt.Println("-------> 1")
		case <-chan2:
			fmt.Println("-------> 2")
		case <-chan3:
			fmt.Println("-------> 3")
		case <-chan4:
			fmt.Println("-------> 4")
		case <-chan5:
			fmt.Println("-------> 5")
		}
		time.Sleep(time.Second * 1)
	}
}
func main() {


	go func1()
	time.Sleep(time.Second * 100)
}
