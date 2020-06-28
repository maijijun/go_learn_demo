package main

import "fmt"

func main() {
	//内置函数

	// len()

	// new(),用来分配内存（值类型）
	num1 := new(int)
	fmt.Printf("%T,%v",num1,num1)

	//make(),用来分配内存（引用类型）

	//cap() 获取容量大小

	//delete 删除，key-value 根据key查找

}
