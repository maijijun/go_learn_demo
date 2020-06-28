package main

import (
	"fmt"
)
func main() {
	name := []string{"Jack","hello","Tom"}
	var nameInput = ""
	fmt.Println("请输入名称：")
	fmt.Scanln(&nameInput)

	//顺序查找
	for	i := 0 ; i < len(name); i++ {
		if name[i] == nameInput{
			fmt.Println("找到名称",name[i],i)
			break
		}else if i == len(name) - 1 {
			fmt.Println("没找到")
		}
	}

	index := -1
	for	i := 0 ; i < len(name); i++ {
		if name[i] == nameInput{
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Println("找到名称",name[index],index)
	}else {
		fmt.Println("没找到")
	}

	fmt.Println("ok")
}