package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//break 跳出当前循环
	//time.Now().Unix() 秒
	//time.Now().UnixNano() 纳秒
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1
	fmt.Println(n)

	for i := 0; i < 1000; i++ {
		if (i == 99) {
			break
		}
		fmt.Println(i)
	}

	//break 可以指定标签 跳出指点 循环
	label1:
	for j := 0; j < 1000; j++  {

		for i := 0; i < 1000; i++ {
			if (i == 99) {
				break label1
			}
			fmt.Println(i)
		}
	}
}
