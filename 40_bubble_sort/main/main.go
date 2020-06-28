package main

import (
	"fmt"
)

func BubbleSort (arr *[]int) {
	var temp = 0
	for j := len(*arr); j > 0; j--{
		for i:= 0; i < j - 1; i++ {
			if (*arr)[i] < (*arr)[i + 1] {
				//交换两者
				temp = (*arr)[i + 1]
				(*arr)[i + 1] = (*arr)[i]
				(*arr)[i] = temp
			}
		}
	}
	fmt.Println(*arr)
}

func main() {
	//内部排序 ，交换式排序，冒泡排序

	intArray := []int{1,22,3,12,223,5,2}
	var temp = 0
	for j := len(intArray); j > 0; j--{
		for i:= 0; i < j - 1; i++ {
			if intArray[i] > intArray[i + 1] {
				//交换两者
				temp = intArray[i + 1]
				intArray[i + 1] = intArray[i]
				intArray[i] = temp
			}
		}
	}
	fmt.Println(intArray)
	BubbleSort(&intArray)
	fmt.Println("ok")
}