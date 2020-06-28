package main

import (
	"fmt"
)

func main() {
	//切片是数组的一个引用
	//大小可以动态改变的
	//slice 底层来说是一个struct，包含（ptr，len，cap）
	var strs [5]string = [...]string{"1", "2", "3", "4", "5"}
	//定义切片
	//	slice := strs[:3]
	//	slice := strs[1:]
	//	slice := strs[:]
	slice := strs[1:3]
	fmt.Println(strs)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//make 创建切片
	//make 创建切片对应的数组对外不可见，只能通过slice访问
	var slice1 []int = make([]int, 6, 12)
	fmt.Println(slice1)

	//直接用数组初始化切片
	var slice2 []int = []int{1, 2, 3}
	fmt.Println(slice2)

	//切片遍历， for，for-range
	for index, val := range slice1 {
		fmt.Println(index, val)
	}

	//切片可以继续切片
	slice3 := slice2[1:2]
	fmt.Println(slice3)

	//切片动态增长
	slice3 = append(slice3, 100, 200)
	fmt.Println(slice3)
	//追加切片
	slice4 := append(slice3, slice1...)
	fmt.Println(slice4)

	//切片拷贝,备份文件相互独立，互不影响
	//拷贝是备份的slice长度不足不会报错，能拷贝多少拷贝多少
	slice6 := make([]int,10,10)
	copy(slice6,slice3)
	fmt.Println(slice6)

	fmt.Println("ok")
}
