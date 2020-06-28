package main

import "fmt"

func main () {
    //go的字符串由字节组成
    //单字符用byte存储
    //编码utf-8

    var a byte = 'a'
    fmt.Println(a)

    fmt.Printf("a = %c\n",a)

    //保存码值大于byte（255）
    var word int =  '北'
    fmt.Printf("汉字是：%c", word)

}
