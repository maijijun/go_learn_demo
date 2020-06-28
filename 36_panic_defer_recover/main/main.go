package main

import (
	"errors"
	"fmt"
)

//panic 抛出异常， recover捕获异常
func test (){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			//自定义错误
			errors.New("读取失败")
		}
	}()
	num1 := 100
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
}

func main() {
	test()
	fmt.Println("ok")

	//抛出错误并终止
	//panic(err)
}
