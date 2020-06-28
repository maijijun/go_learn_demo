package main

import (
	"fmt"
)


//defer 延时机制
//当执行到defer，defer后面的语句压入栈中，不执行
//当函数执行完毕后，在按照先入后出的方式执行
//资源释放的时候经常使用
//有值操作时，会先将值拷贝入栈
func deferTest(){
	defer fmt.Println("ok1")
	fmt.Println("ok2")
	defer fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	n := 1
	n++
	defer fmt.Println(n)
	n++
	fmt.Println(n)
}
func main() {

	deferTest()

	fmt.Println("okn")
}
