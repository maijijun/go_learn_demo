package main

import (
	"fmt"
)

func main() {
	var i int = 100
	fmt.Println("i在内存中的地址",&i)

	var ptr *int = &i
	fmt.Printf("ptr的值 = %v\n",ptr)
	fmt.Printf("ptr的地址 = %v\n",&ptr)
	fmt.Printf("ptr指向的值 = %v\n", *ptr)

}
