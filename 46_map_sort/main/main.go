package main

import (
	"fmt"
	"sort"
)

type Stu struct{
	Name string
	Age int
}
func main() {
	map1 :=  map[int]string{
		1:"1",
		2:"2",
		3:"221",
		4:"12",
		5:"112",
	}
	fmt.Println(map1)
	//map 排序
	//将map的 key 放入切片
	//排序切片
	var keys []int
	for key,_ := range map1 {
		keys = append(keys, key)
	}
	fmt.Println("keys:",keys)
	sort.Ints(keys)
	for _, key := range keys{
		fmt.Println(map1[key])
	}

	//map的值一般用struct

	students := make(map[int]Stu, 10)
	stu1 := Stu{"Tom",12}
	stu2 := Stu{"Tom1",112}
	students[1111] =stu1
	students[12121] =stu2
	fmt.Println(students)


	fmt.Println("ok")
}
