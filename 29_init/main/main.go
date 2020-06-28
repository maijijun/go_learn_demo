package main

import "fmt"

var a = test2()

func test2() int {
	fmt.Println("全局变量")
	return 90
}

func test(a int) {
	if a > 2 {
		a--
		test(a)
	}
	fmt.Println("1---", a)
}
//每个函数都可以包含一个 init 函数
//init 在 main 之前执行
//做初始化操作
//全局变量定义 -> init -> main
func init()  {
	fmt.Println("执行 init")
}
func main() {
	test(5)
}
