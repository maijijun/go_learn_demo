package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct{
	Name string
	Age int
}

type Teacher struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//go 没有class的概念，go支持面向对象的编程呢个， 面向对象的概念使用 结构体 实现， 对象
	//struct 是值类型

	//定义1
	var stu1  Stu
	fmt.Println(stu1)

	//赋值
	stu1.Name = "Jack"
	stu1.Age = 10
	fmt.Println(stu1.Name , "  ", stu1.Age)

	//定义2
	stu2 := Stu{"a",13}
	fmt.Println(stu2)

	//定义3
	var stu3 *Stu = new(Stu)
	(*stu3).Age = 12
	stu3.Name ="b"
	fmt.Println(*stu3)

	//定义3
	var stu4 *Stu = &Stu{}
	stu4.Name = "c"
	(*stu4).Age = 111
	fmt.Println(stu4)

	//结构体的所有字段是连续分布的

	//结构体转换要求字段 个数 名称 类型 相同

	//结构体可以使用 tag ，可以通过反射机制来获取，常见的是序列化和反序列化
	//序列化Json格式
	var teacher Teacher
	teacher.Name ="Jack"
	teacher.Age = 100
	teacher.Skill = "ha~~~me~~~ha~~me~~ha~~~"
	teacherJson,err := json.Marshal(teacher)
	fmt.Println(string(teacherJson),"err:",err)

	fmt.Println("ok")
}
