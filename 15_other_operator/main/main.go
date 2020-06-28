package main

import "fmt"

func main() {
	// &(取地址)  *（取值）
	var a = 100
	var b *int = &a
	fmt.Println(a,b,*b,&a)
}
