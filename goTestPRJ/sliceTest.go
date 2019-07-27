package main

import "fmt"

var num int =0

func main() {
num := 1
num++
var slice1 = []int{1, 2, 3, 4, 5, 6}
lencap(slice1)

slice2 := slice1[0:4:4]
lencap(slice2)

var a [5]int = [5]int{1, 2, 3, 4, 5} //先定义了一个数组
array_slice := a[:]
fmt.Printf("cap:%d a_addr:%p slice_addr:%p slice_type:%T array_type:%T\n", cap(array_slice), &a, &array_slice[0], array_slice, a)
lencap(array_slice)

array_slice[1] = 9
fmt.Printf("cap:%d addr:%p value:%v  slice:%v\n", cap(array_slice), &array_slice[0], a, array_slice)
lencap(array_slice)


array_slice = append(array_slice, 6)
fmt.Printf("cap:%d addr:%p value:%v\n", cap(array_slice), &array_slice[0], array_slice)
lencap(array_slice)

array_slice = append(array_slice, 7)
fmt.Printf("cap:%d addr:%p value:%v\n", cap(array_slice), &array_slice[0], array_slice)
lencap(array_slice)

f(array_slice)
fmt.Printf("cap:%d addr:%p value:%v\n", cap(array_slice), &array_slice[0], array_slice)
lencap(array_slice)

array_slice = array_slice[:8]
fmt.Printf("-----8 cap:%d addr:%p value:%v\n", cap(array_slice), &array_slice[0], array_slice)
lencap(array_slice)

var s string = "abcdefg"
string_slice := s[0:5]
fmt.Printf("%p %T\n", &s, string_slice)
array_slice = append(array_slice, 1)
array_slice = append(array_slice, 1)
array_slice = append(array_slice, 1)
fmt.Printf("-----8 cap:%d addr:%p value:%v\n", cap(array_slice), &array_slice[0], array_slice)

}
func f(a []int) {
	a = append(a, 8)
	fmt.Printf("f cap:%d addr:%p value:%v\n", cap(a), &a[0], a)
}

func lencap(a []int) {
	num++
	fmt.Printf("------->%d	%v %d %d\n\n", num, a, len(a), cap(a))
}
