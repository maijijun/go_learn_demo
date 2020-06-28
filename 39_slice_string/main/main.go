package main

import (
	"fmt"
)

func main() {
	//string 底层是byte数组，可以进行切片处理
	str := "hello@123As"
	fmt.Println(str)
	slice :=str[1:4]
	fmt.Println(slice)

	fmt.Println("ok")
}
