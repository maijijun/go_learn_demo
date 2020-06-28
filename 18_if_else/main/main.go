package main

import "fmt"

func main() {

	//if, else,  else if
	var age int = 0
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if (age > 20) {
		fmt.Println("你需要负责")
	}else {
		fmt.Println("你滚吧")
	}

}
