package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "hello"
	//len计算长度
	//汉字占 3个字节， 字母占 1个字节
	fmt.Println(len(str))

	//Atoi 字符串转整数
	n, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("转换失败")
	}
	fmt.Println(n)

	//Itoa 整数转字符串
	a := strconv.Itoa(1)
	fmt.Println(a)

	//字符串转 []byte
	bytes := []byte("hello")
	fmt.Println(bytes)

	//[]byte 转字符串
	str1 := string([]byte{97, 98, 99})
	fmt.Println(str1)

	//FormatInt 10进制转其他

	//查找字符串中是否包含
	fmt.Println(strings.Contains("abcde", "ab"))

	//统计字符串中的指定字符串的个数
	fmt.Println(strings.Count("aaaaaaabaaabbsaa", "b"))

	//比较字符串  ==（区分大小写）  或 EqualFold（不区分大小写）
	fmt.Println(strings.EqualFold("aaa", "aaa"))

	//子串第一次出现的下标
	fmt.Println(strings.Index("aaasasas", "s"))

	//子串最后一次出现的下标
	fmt.Println(strings.LastIndex("aaasasas", "s"))

	//子串替换 n = -1 全部替换，n表示希望替换几个
	fmt.Println(strings.Replace("aasasasasasa", "a", "GO", -1))

	//拆分字符串
	fmt.Println(strings.Split("de,dsd,eda,dw", ","))

	//大小写转换
	fmt.Println(strings.ToLower("aaAAASA"))
	fmt.Println(strings.ToUpper("aaaaAAAAS"))

	//去掉两边空格 , 去掉左右指定字符 Trim("hell0!","!")  去左TrimLeft 去右 TrimRight
	fmt.Println(strings.TrimSpace("   sasas   "))

	//HasSuffix   HHasPrefix  是否有什么开头/结尾
}
