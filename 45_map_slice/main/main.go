package main

import "fmt"

func main() {
	//map 切片
	//map的个数可以动态变化

	var mapSlice []map[int]string
	//make  切片
	mapSlice = make([]map[int]string, 2)
	//make map
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[int]string, 2)
		mapSlice[0][1] = "Jack"
		mapSlice[0][2] = "Tom"
	}
	if mapSlice[1] == nil {
		mapSlice[1] = make(map[int]string, 2)
		mapSlice[1][1] = "Julia"
		mapSlice[1][2] = "Back"
	}
	fmt.Println(mapSlice)

	//append 动态增加
	mapn := map[int]string{
		1:"map",
		2:"Tim",
	}
	mapSlice = append(mapSlice, mapn)
	fmt.Println(mapSlice)
	fmt.Println("ok")
}
