package main

import (
	"fmt"
	"strings"
)

//累加器
//闭包
//函数与变量构成一个整体
func AddUpper() func(int) int {
	var n = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	//闭包 和外面的suffix
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
	//闭包
}
func main() {

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	//返回一个闭包
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("name"))
	fmt.Println(f1("name1.jpg"))
}
