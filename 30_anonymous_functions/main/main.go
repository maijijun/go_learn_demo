package main

import "fmt"

func main() {
	n1 := 1
	n2 := 2
	//在定义匿名函数时直接调用
	res1 := func(n1, n2 int) int {
		return n1 + n2
	} (n1, n2)

	fmt.Println(res1)

	//将匿名函数赋值给变量

	a := func(n1, n2 int) int {
		return n1 + n2
	}
	res2 := a(10, 9)
	fmt.Println(res2)
}
