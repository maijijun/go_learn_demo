package main

import (
	"fmt"
	"go_learn_demo/51_encapsulation/model"
	model2 "go_learn_demo/52_extends/model"
)

func main() {
	//继承
	//减少代码冗余
	//不利于代码维护和功能扩展
	// go 没有extends 关键字
	//go 用匿名结构体实现继承

	teacher := model.NewTeacher()
	teacher.SetName("Jam")
	teacher.SetAge(10)
	fmt.Println(*teacher)

	bigTeacher := model2.NewBigTeacher()
	bigTeacher.Teacher.SetName("Jack")
	bigTeacher.Teacher.SetAge(21)
	bigTeacher.SetSkill("say hello")
	fmt.Println(*bigTeacher)

	//可以简化
	fmt.Println(bigTeacher.Teacher.GetAge()," ",bigTeacher.GetName())

	//多重继承
	//访问 有名结构体 时必须带上名称
	//本身不带 name ，结构体 A 和 B 带 name ，赋值时需要指定

	// 结构体可以嵌入 匿名基本数据类型，同类型最多只有一个
	fmt.Println("ok")
}
