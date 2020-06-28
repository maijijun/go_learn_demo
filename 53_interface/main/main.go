package main

import (
	"fmt"
)

//定义一个接口
type Usb interface {
	Start()
	Stop()
	Say
}

type Say interface {
	Hello()
}

type Phone struct {

}

func (p Phone) Start()  {
	fmt.Println("Phone Start")
}

func (p Phone) Stop()  {
	fmt.Println("Phone Stop")
}

func (p Phone) Hello()  {
	fmt.Println("Phone Say")
}

type Camera struct {

}

func (c Camera) Start()  {
	fmt.Println("Camera Start")
}

func (c Camera) Stop()  {
	fmt.Println("Camera Stop")
}

func (c Camera) Hello()  {
	fmt.Println("Camera Say Hello")
}
//定义一个computer 实现接口Usb

type Computer struct {

}

func (c Computer) working(usb Usb)  {
	usb.Start()
	usb.Stop()
	usb.Hello()
}

type NullI interface {

}




func main() {

	//接口不能包含任何变量
	//接口中可以定义方法
	//接口中的方法不能有实现体
	//高内聚低耦合
	//多态
	//go 中不需要显式实现接口
	// go 没有 implement 关键字
	//接口可以对结构体进行扩展，不破坏器原本特性
	cp := Computer{}
	ph := Phone{}
	ca := Camera{}

	cp.working(ph)
	cp.working(ca)

	//Camera 实现了 Usb 接口
	var c Camera
	var u Usb = c
	u.Start()

	//可以实现多个接口
	var c1 Camera
	var u1 Usb = c1
	var s1 Say = c1
	u1.Stop()
	s1.Hello()
	//接口可以继承，继承需要实现继承接口的函数
	u1.Hello()

	//空接口没有任何方法，所以所有类型实现了空接口，我们可以把任何一个变量给空接口
	//interface{}

	 str := "hello"
	var null1 NullI = str
	fmt.Println(null1)

	fmt.Println("ok")
}
