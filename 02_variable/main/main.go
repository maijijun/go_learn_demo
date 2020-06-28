package main

import "fmt"

func main () {
	//1.指定变量类型
	var i int
	fmt.Println("int = " , i)

	//2.根据值自动判断类型
	var num = 100.11
	var name = "Jack"
	var age = 10
	fmt.Println("num = ", num, "name = ", name, "age = ",age)

	//3.用 := 声明并赋值，省略 var
	title := "天之痕"

	fmt.Println("title = ", title)

	//4.一次性声明多个变量
	var a, b, c int
	fmt.Println(a,b,c)

	n, m, l := "name", 2, 100.1
	fmt.Println(n, m, l)

	var (
		s = 10
		d = "111"
	)
	fmt.Println(s,d)
}
