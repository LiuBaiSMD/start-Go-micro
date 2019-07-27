package main

import (
	"fmt"
	"time"
)

//import (
//	"time"
//	"fmt"
//)
//
//func pub() {
//	tick := time.NewTicker(time.Second)
//	i := 0
//	for range tick.C {
//		fmt.Println(time.Now().String())
//		//fmt.Println(1)
//
//		i++
//	}
//}
//
//func main(){
//	pub()
//}


func main() {
	//closeChannel()
	c := make(chan int)
	timeout := time.After(time.Second * 2) //
	t1 := time.NewTimer(time.Second * 3)   // 效果相同 只执行一次
	var i int
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C:      // 代码段2
				fmt.Println("3s定时任务")
			case <-timeout:   // 代码段1
				i++
				fmt.Println(i, "2s定时输出")
			case <-time.After(time.Second * 4):    // 代码段3
				fmt.Println("4s timeout。。。。")
			default:                               // 代码段4
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 600)
	//close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}