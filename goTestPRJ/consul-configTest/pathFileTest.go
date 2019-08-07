package main

import (
	"fmt"
	"path/filepath"
)

func main(){
	data1 := string(filepath.Separator)
	data2 := filepath.Join("./", string(filepath.Separator))
	data3 := filepath.Dir(filepath.Join("./", string(filepath.Separator)))
	data4,_ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))
	fmt.Println("----->", data1)
	fmt.Println("----->", data2)
	fmt.Println("----->", data3)
	fmt.Println("----->", data4)
}
