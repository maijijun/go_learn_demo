package main

import "fmt"

func main() {
	var str string = "hello 北京"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c",str2[i])
	}

	for index ,val := range str {
			fmt.Printf("%d %c",index,  val)
	}

}
