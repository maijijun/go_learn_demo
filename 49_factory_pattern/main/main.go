package main

import (
	"fmt"
	"go_learn_demo/49_factory_pattern/model"
)

func main() {
	//go 用工厂模式来实现构造函数

	student := model.Student{
		Name: "Jack",
		Age:  0,
	}
	fmt.Println(student)

	//工厂模式 teacher 结构体
	teacher := model.NewTeacher("Tom",33)
	fmt.Println(*teacher," ",teacher.GetAge())
	fmt.Println("ok")
}
