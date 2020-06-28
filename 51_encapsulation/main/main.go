package main

import (
	"fmt"
	"go_learn_demo/51_encapsulation/model"
)

func main() {
	//封装
	//把抽象出来的字段和对字段的操作封装在一起
	//隐藏细节
	//数据校验
	//实现，结构体，字段 小写

	teacher := model.NewTeacher()
	teacher.SetName("Jam")
	teacher.SetAge(10)
	fmt.Println(*teacher)

	fmt.Println("ok")
}
