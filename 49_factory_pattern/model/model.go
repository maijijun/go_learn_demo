package model

type Student struct {
	Name string
	Age int
}

//首字母小写不能引入，可以用工厂模式解决
type teacher struct {
	Name string
	age int
}

//定义构造函数
func NewTeacher(name string, age int )  *teacher {
	newTeacher := teacher{
		Name: name,
		age:  age,
	}
	return &newTeacher
}

//定义get方法获取age
func (t *teacher) GetAge() int{
	return t.age
}
