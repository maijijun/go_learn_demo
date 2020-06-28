package model

import "fmt"

type teacher struct {
	name string
	age int
}

func NewTeacher() *teacher {
		return &teacher{
		}
}

func (teacher *teacher) GetName() string {
	return teacher.name
}

func (teacher *teacher) SetName(name string) {
	teacher.name = name
}

func (teacher *teacher) GetAge() int {
	return teacher.age
}

func (teacher *teacher) SetAge(age int) {
	if age > 150 || age < 0 {
		fmt.Println("输入年龄异常")
	}else {
		teacher.age = age
	}
}
