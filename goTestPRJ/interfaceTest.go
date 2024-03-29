package main

import "fmt"

type S struct {
	i int
	j string
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(int)
}

func f1(p I) {
	fmt.Println(p.Get())
	p.Put(888)
}

func f2(p ...interface{}) {
	switch t := p[0].(type) {
	case int:
		fmt.Println("this is int number")
	case I:
		fmt.Println(t)
		fmt.Println("I:", t.Get())
	default:
		fmt.Println("unknow type")
	}
}

//指针修改原数据
func dd(a *S) {
	a.Put(999)
	fmt.Println(a.Get(), "in dd func")
}

//临时数据
func aa(a S) {
	a.Put(2222)
	fmt.Println(a.Get(), "in aa func")
}

func main() {
	var s S
	s.i = 1
	s.j = "wuxun"
	//s.Put(333)
	//fmt.Println(s.Get())
	//f1(&s)
	fmt.Println(s.Get())
	fmt.Println(s.i)
	f2("hello","test")
	//dd(&s)
	//fmt.Println(s.Get())
	//aa(s)
	//fmt.Println(s.Get())

}