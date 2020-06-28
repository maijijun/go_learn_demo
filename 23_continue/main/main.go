package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//continue 结束本次循环

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1
	fmt.Println(n)

	for i := 0; i < 1000; i++ {
		if (i == 99) {
			continue
		}
		fmt.Println(i)
	}

	//break 可以指定标签 跳出指点 循环
	label1:
	for j := 0; j < 1000; j++  {

		for i := 0; i < 1000; i++ {
			if (i == 99) {
				continue label1
			}
			fmt.Println(i)
		}
	}
}
