package main

import "fmt"

func main() {
	//goto 尽量不要使用，容易混乱

	fmt.Println("ok1")
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	label1:
	fmt.Println("ok5")
	fmt.Println("ok6")

}
