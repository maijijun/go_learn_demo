package main

import (
	"fmt"
)

func BubbleSort (arr *[]int) {
	var temp = 0
	for j := len(*arr); j > 0; j--{
		for i:= 0; i < j - 1; i++ {
			if (*arr)[i] > (*arr)[i + 1] {
				//交换两者
				temp = (*arr)[i + 1]
				(*arr)[i + 1] = (*arr)[i]
				(*arr)[i] = temp
			}
		}
	}
	fmt.Println(*arr)
}

func BinaryFind (arr *[]int, leftIndex int, rightIndex int, findVal int) {
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	//有序数组升序
	if findVal < (*arr)[middle] {
		//找左边
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if  findVal > (*arr)[middle] {
		//找右边
		BinaryFind(arr, middle+1, rightIndex, findVal)
	}else {
		fmt.Println("找到了",(*arr)[middle], middle)
	}
}

func main() {
	//二分查找
	//有序数组
	//找出中间值判断
	//递归
	name := []int{1,2,3,55,44,8,90,123,1211,0}
	num := 0
	fmt.Println("请输入数字")
	fmt.Scanln(&num)
	//升序
	BubbleSort(&name)
	BinaryFind(&name,0, len(name), num)

	fmt.Println("ok")
}