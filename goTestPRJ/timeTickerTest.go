package main

import (
	"time"
	"fmt"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		fmt.Println(time.Now().String())
		//fmt.Println(1)

		i++
	}
}

func main(){
	pub()
}