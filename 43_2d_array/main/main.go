package main

import (
	"fmt"
)
func main() {
	//定义和一维数组类似
	var arr [3][2]int = [3][2]int{{1,2},{1,1},{1,2}}
	fmt.Println(arr)

	//遍历， 双重for循环， for-range
	for index, val := range arr {
		fmt.Println(index,"=",val)
		for index,val := range val {
			fmt.Println(index , "=", val)
		}
	}
}