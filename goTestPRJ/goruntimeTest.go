package main

import (
	"fmt"
	"time"
)

func func1(i int){
	for{
		fmt.Println(i)
		time.Sleep(time.Second)
	}

}

func main() {
	i := 1
	for {
		if i==10{
			break
		}
		i++
		go func1(i)
		time.Sleep(time.Second)
		defer func1(i)
	}
	time.Sleep(time.Second * 100)
	fmt.Println("------> over")
}