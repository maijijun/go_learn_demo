package model

import "fmt"

type Teacher struct {
	name string
	age  int
}

func NewTeacher() *Teacher {
	return &Teacher{
	}
}

func (teacher *Teacher) GetName() string {
	return teacher.name
}

func (teacher *Teacher) SetName(name string) {
	teacher.name = name
}

func (teacher *Teacher) GetAge() int {
	return teacher.age
}

func (teacher *Teacher) SetAge(age int) {
	if age > 150 || age < 0 {
		fmt.Println("输入年龄异常")
	} else {
		teacher.age = age
	}
}

type bigTeacher struct {
	Teacher //嵌入了匿名结构体（不含变量名）
	skill   string
}

func NewBigTeacher() *bigTeacher {
	return &bigTeacher{
	}
}

func (bigTeacher *bigTeacher) GetSkill() string {
	return bigTeacher.skill
}

func (bigTeacher *bigTeacher) SetSkill(skill string) {
	bigTeacher.skill = skill
}
