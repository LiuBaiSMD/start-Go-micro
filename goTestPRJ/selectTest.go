package main

import (
	"fmt"
	"time"
)

var ch1 chan int
var ch2 chan int
//如果在此处声明则没有分配地址空间，将不起任何作用
//var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func getNumber(i int) int {
fmt.Printf("numbers[%d]\n", i)

return numbers[i]
}
func getChan(i int, chs []chan int) chan int {
fmt.Printf("chs[%d]\n", i)
fmt.Println(&chs, &chs[i])
return chs[i]
}
func main () {
	ch1 =make(chan int, 1)
	ch2 =make(chan int, 1)
	chs := []chan int{ch1, ch2}
	fmt.Println(&chs, &ch1, &ch2)
	go func(){
		for{
			fmt.Println("---->start send")
			time.Sleep(time.Second)
			fmt.Println("---->sleep over")
			fmt.Println(len(ch1), len(ch2), cap(ch1), cap(ch2))
			<-ch1
			<-ch2
			fmt.Println("---->send over")
		}
	}()

	go func(){
		for{
			time.Sleep(time.Second)
			select {
			case getChan(0, chs) <- getNumber(2):
				fmt.Println("1th case is selected.")
			case getChan(1, chs) <- getNumber(3):
				fmt.Println("2th case is selected.")
			default:
				fmt.Println("default!.\n")
			}
			fmt.Println()
		}

	}()

	time.Sleep(time.Second)
	//fmt.Println(len(ch1), len(ch2))
	if len(ch1)>0{
		num1 := <- ch1
		fmt.Println("chan1:		", num1)

	}
	if len(ch2)>0{
		num2 := <-ch2
		fmt.Println("chan2		", num2)
	}
	time.Sleep(time.Second * 100)
}