package main

import (
	"time"
	"fmt"
)



func main()  {
	tchan := time.NewTimer(time.Second*3)
	tchan1 := time.After(time.Second*3)
	tchan2 := time.NewTicker(time.Second * 3)
	fmt.Printf("tchan type=%T\n",tchan, tchan)
	fmt.Println("mark 1")
	//fmt.Println("tchan=",<-tchan)
	fmt.Println("mark 2")
	for {
		select {
		case <-time.After(time.Second * 2):
			time.Sleep(time.Second * 2)
			fmt.Println("1------->")
		case <-tchan.C:
			fmt.Println("2------>NewTimer")
		case <-tchan1:
			fmt.Println("3------>After")
		case <-tchan2.C:
			fmt.Println("4------>Ticker")
		}
	}

	fmt.Println("over")
}