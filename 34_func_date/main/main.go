package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前日期
	now := time.Now()
	fmt.Println(now)

	//去年月日 now.Year ,now.Day ....

	//时间常量
	fmt.Println(time.Second)
	fmt.Println(time.Nanosecond)

	//Unix UnixNao
	fmt.Println(now.Unix(), now.UnixNano())
}
