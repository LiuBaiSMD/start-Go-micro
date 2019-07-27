package main

import(
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rad := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 20; i++ {
		fmt.Println(rad.Intn(100))
	}
	for i := 0; i < 20; i++ {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println("---------->")
	for i := 0; i < 20; i++ {
		fmt.Println(rand.Intn(100))
	}
}
