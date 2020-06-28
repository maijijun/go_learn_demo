package main

import "fmt"

func main() {
	//默认 ""
    var address string = "北京"
    fmt.Println("address = ", address)

    //字符串赋值后不可变
    //var str = "aaaa"
    //str[0] = 'b'

    // 漂符号  ``  转义 字符串

    var str = `http://123.com` + `121212121`
    fmt.Println(str)

    //字符串 拼接 换行

    var str1 = "11212" +
    	"1212121" +
    	"1212121"
    fmt.Println(str1)
}
