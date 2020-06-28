package main

import (
	"fmt"
)

func main() {
	//数组定义后长度不能改变
	var hens [4]float64
	hens[0] = 1.1
	hens[1] = 12.1
	hens[2] = 11.1
	hens[3] = 4.1
	var sum float64
	for i :=0; i < len(hens); i++ {
		sum += hens[i]
	}
	fmt.Println(sum)

	var h [3]int = [3]int{1,2,3}
	fmt.Println(h)

	var h1 = [3]int{1,2,3}
	fmt.Println(h1)

	var h2 = [...]int{1,2,3}
	fmt.Println(h2)

	var h3 = [...]int{1:100,2:90,3:0,0:1212}
	fmt.Println(h3)

	fmt.Println("ok")
}
