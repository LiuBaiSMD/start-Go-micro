package main

import (
	"fmt"
	"time"
)

func func1( chan1 chan string ){
	data := ""
	for{
//select {
		//case data=<-chan1:
		data=<-chan1
				data=data+"1 "
				fmt.Println(data)
				time.Sleep(time.Second * 1)
				chan1<-data
			//}
		}

	}
func func2( chan1 chan string ){
	data := ""
	for{
//select {
		//case data=<-chan1:
		data=<-chan1
			data=data+"2 "
			fmt.Println(data)
			time.Sleep(time.Second * 1)
		chan1<-data
		//}
	}

}

func func3( chan1 chan string ){
	data := ""
	for{
		//select {
		//case data=<-chan1:
		data=<-chan1
			data=data+"3 "
			fmt.Println(data)
			time.Sleep(time.Second * 1)
			chan1<-data
		//}
	}

}
func func4( chan1 chan string ){
	data := ""
	for{
//select {
		//case data=<-chan1:
		data=<-chan1
			data=data+"4 "
			fmt.Println(data)
			time.Sleep(time.Second * 1)
			chan1<-data
		//}
	}

}

var chan1 chan string

func main() {
	chan1 = make(chan string )
	go func1(chan1)
	go func2(chan1)
	go func3(chan1)
	go func4(chan1)
	time.Sleep(time.Second * 2)
	chan1<-"0 "
	time.Sleep(time.Second * 100)
}
