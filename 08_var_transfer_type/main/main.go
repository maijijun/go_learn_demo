package main

import (
	"fmt"
	"strconv"
)

func main() {
	//go 基本数据类型转换
	// %d int, %f float, %t bool, %c char

	var i int = 100

	var j float32 = float32(i)

	var str string  = ""
	fmt.Println(i,j)

	//基本数据类型转 string
	//1. fmt.Sprintf (推荐)
	var num1 int = 100
	str = fmt.Sprintf("num1 =%d", num1)
	fmt.Println(str)

	//2. strconv
	str = strconv.FormatInt(int64(num1),2)
	fmt.Println(str)

	var num3 int = 11111
	str = strconv.Itoa(num3)
	fmt.Println(str)

	//string 转 基本数据类型, 无效数据变为 0

	var f string = "true"
	var b bool
	b,_ = strconv.ParseBool(f)
	fmt.Printf("%T %v",b, b)

}
