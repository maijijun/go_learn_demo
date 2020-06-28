package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//定义 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

//定义 hero 结构体 切片
type HeroSlice  []Hero


//实现 Sort 接口 方法
//长度
func (hs HeroSlice) Len() int{
	return len(hs)
}
//排序条件 ,age
func (hs HeroSlice) Less(i, j int) bool{
	return hs[i].Age > hs[j].Age
}
//交换
func (hs HeroSlice) Swap(i, j int)  {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i],hs[j] = hs[j],hs[i]
}

func main() {

	//对结构体切片排序

	var heroes HeroSlice
	for i := 0; i < 10; i++{
		hero := Hero{
			Name: fmt.Sprintf("Hero~%d",rand.Intn(100)),
			Age:  rand.Intn(80),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println("排序前~~~")
	for _,val := range heroes {
		fmt.Println(val)
	}
	fmt.Println("排序后~~~")
	sort.Sort(heroes)
	for _,val := range heroes {
		fmt.Println(val)
	}
	fmt.Println("ok")
}
