package main

import "fmt"

func main() {
	// &&(短路与) ||（短路或） !
	i := 1
	if i < 2 {
		fmt.Println(i)
	}
}
