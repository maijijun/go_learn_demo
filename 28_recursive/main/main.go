package main

import "fmt"

func test(a int) {
	if a > 2 {
		a--
		test(a)
	}
	fmt.Println("1---", a)
}

func test2(a int) {
	if a > 2 {
		a--
		test2(a)
	}else {
		fmt.Println("2---", a)
	}
}
func main() {
	test(5)
	test2(5)
}
