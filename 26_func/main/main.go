package main

import "fmt"

func cal (a int, b int) int{
	c := a - b
	return c
}
//函数也是一种数据类型
//go不支持重载
func main() {
	fmt.Println(cal(3,1))

	//给 int 取别名
	type myInt int

	var i myInt
	fmt.Printf("%T %v",i,i)
}


