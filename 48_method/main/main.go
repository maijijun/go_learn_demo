package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}
//该方法只能 Stu 结构体使用
func (stu Stu) say() {
	fmt.Println(stu.Name, "Say Hello")
}

type Teacher struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (teacher *Teacher) say() {
	fmt.Printf(teacher.Name, "Say Hello")
}

func (teacher *Teacher) String() string {
	str := fmt.Sprintf("Name= [%v],Age = [%v]",teacher.Name,teacher.Age)
	return str
}

type Integer int

func (integer *Integer) print() {
	fmt.Println("n =", *integer)
}

func main() {
	//go 方法是作用在指定的数据类型上的，不止是struct可以使用 （type 定义）

	//值拷贝
	var stu1 Stu
	stu1.Name = "Jack"
	stu1.say()

	//通常方法和结构体的指针类型绑定(地址传递)
	var teacher1 Teacher
	(&teacher1).Name = "Tom"
	teacher1.say()

	//int float32等也可以有自己的方法

	var i Integer
	i = 1000
	i.print()

	//如果实现 String() 方法，fmt.Println() 默认调用 String()
	var teacher2 Teacher
	teacher2.Name = "Mrs"
	teacher2.Age = 2
	fmt.Println(&teacher2)

	//方法和函数的区别
	//调用方式不同
	//函数接收者为值类型，不能将指针直接传递，方法可以

	var t = Teacher{
		Name:  "Julia",
		Age:   0,
		Skill: "Kung fu",
	}
	fmt.Println(t)

	t2 := &Teacher{
		Name:  "P",
		Age:   0,
		Skill: "p",
	}
	fmt.Println(t2)

	fmt.Println("ok")
}
