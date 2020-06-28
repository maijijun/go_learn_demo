package main

import (
	"fmt"
	"unsafe"
)
func main () {
	//默认int 0
	// int8 -128~127
	//uint8 0~255
	//int16, int32, int64
	//uint16, uint32, uint64
	//byte 0~255

	//查看变量的类型
	var n1 int8
	fmt.Printf("n1的数据类型 %T\n",n1)

	//查看变量的占用字节大小
	fmt.Printf("n1占用的字节大小 %d", unsafe.Sizeof(n1))
}
