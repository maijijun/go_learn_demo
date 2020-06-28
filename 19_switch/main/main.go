package main

import "fmt"

func main() {

	//go switch 不需要 break， 默认会加

	var age int = 0
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	//age.(type) 获取接口类型
	switch(age) {
	case 1:
		fmt.Println("1")
	case 2:
		//switch穿透  fallthrough
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println("滚")

		
	}

}
