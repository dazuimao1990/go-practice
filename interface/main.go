package main

import "fmt"

/*
电脑具备一个 Usb 接口，手机插入的时候被识别为手机相关操作，相机插入时被识别为相机相关操作
接口是实现多态的方式，体现了高内聚低耦合的编程思想
接口在 Golang 开发中大量存在，需要掌握
实现一个接口，就意味着实现了接口中定义的所有方法，方法必须完全一样定义，包括名字、传参、返回
在 Golang 中，只要结构体实现了接口中的所有方法，就认为是实现了接口，不需要显式指定实现关系
接口中定义的只能 是空方法，具体实现是要到结构体那里的方法实现的
*/

// 定义一个接口
type Usb interface {
	// 接口中包含有方法，作为简单示例，没有传参或返回
	Start()
	Stop()
}

type HDMI interface {
	Output()
}

/*
接口可以继承，Enter 继承了上述两个接口，类型在实现接口的时候，必须实现所有的接口方法
包括 Usb 和 HDMI 中定义的
继承的父接口之间如果有相同的方法名，则必须函数签名（函数签名包括方法的名字、参数列表、返回值列表）必须一致，否则会编译报错
*/
type Enter interface {
	Usb
	HDMI
	check()
}

// 定义一个Phone结构体，主要关注方法
type Phone struct {
}

// 在 Phone 中定义方法来实现接口 Usb
func (p *Phone) Start() {
	fmt.Println("手机开始工作。。。")
}

func (p *Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

// 定义一个Camera结构体，主要关注方法
type Camera struct {
}

// 在 Camera 中定义方法来实现接口 Usb
func (p *Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (p *Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

// 定义结构体 Computer 作为使用接口的一方
type Computer struct {
}

// 为 Computer 定义接口调用方法，在这里，将实现了接口 Usb 的结构体作为参数传递
func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var phone Phone
	var camera Camera
	var computer Computer
	var apple Phone
	var a Usb = &apple

	a.Start()

	computer.Working(&phone) // 将 Phone 类型传入，则按照 phone.Start phone.Stop 处理

	computer.Working(&camera) // 将 Camera 类型传入，则按照 camera.Start camera.Stop 处理
}
