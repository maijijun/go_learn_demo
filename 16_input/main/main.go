package main

import "fmt"

func main() {
	// Scanln
	var name string
	var age int
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println(name)

	//Scanf
	fmt.Println("请输入姓名，年龄")
	fmt.Scanf("%s %d",&name,&age)
	fmt.Println(name,age)
}
