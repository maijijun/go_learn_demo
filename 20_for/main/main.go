package main

import "fmt"

func main() {
	var str string = "你好"
	for i := 1; i <= 10; i++ {
		fmt.Println(str)
	}
	//死循环
	a := 0
	for {
		if (a == 2) {
			break
		}
		fmt.Println(a)
		a++
	}
}
