package main

import "fmt"

func main() {
	//map
	//引用类型，在函数中使用会改变原值
	//key-value
	//map[key]value
	//map声明后需要make分配内存
	// key 不能重复
	//go 的 map 是无序的

	//map 定义 使用
	//1
	var a map[int]string
	a = make(map[int]string, 10)
	a[0] ="Jack"
	fmt.Println(a)

	//2
	cities := make(map[int]string)
	cities[0] = "1"
	fmt.Println(cities)

	//3
	map2 :=  map[int]string{
		1:"1",
		2:"2",
	}
	fmt.Println(map2)

	fmt.Println("增删改查")
	//增加
	map2[3] = "4"
	fmt.Println(map2)
	//改
	map2[1] = "3"
	fmt.Println(map2)

	//删
	//删除所有， 遍历key删除， 或make一个新的map，让gc回收
	delete(map2,3)
	fmt.Println(map2)

	//查
	val, findRes := map2[2]
	fmt.Println(val," ",findRes)

	fmt.Println("遍历")
	//遍历
	for index, val := range map2 {
		fmt.Println(index, "=", val)
	}

	//map长度
	fmt.Println(len(map2))

	fmt.Println("ok")
}