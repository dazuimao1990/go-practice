package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

// 方法里会通过 (p Person) 绑定给指定的数据类型，这里不仅仅指 struct 这种类型。
// p 相当于是哪个 Person 类型的变量调用了这个方法，就传谁的值，像传参一样，类似于 Python 里的 self
// P 为数值类型，则进行数值拷贝，如果是引用类型，就进行地址拷贝
func (p Person) speak() {
	fmt.Printf("%v 是个好人\n", p.Name)
}

// 方法可以定义输出
func (p Person) jisuan() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	return sum
}

// 方法可以定义传参
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Printf("%v计算从1到%v的结果是%v.\n", p.Name, n, res)
}

// 方法可以传多个参数，定义多个返回
func (p Person) GetRes(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// 案例：定义结构体 Circle ，字段为半径 radius，声明一个方法 area 和 Circle 绑定，可以返回面积。

type Circle struct {
	Radius float64
}

func (c Circle) area() (res float64) {
	pi := math.Pi
	return pi * c.Radius * c.Radius
}

func main() {
	man := Person{"Tom", 10}
	c := Circle{10}
	man.speak()
	fmt.Println("1+···+100 = ", man.jisuan())
	man.jisuan2(101)
	sum, sub := man.GetRes(21, 12)
	fmt.Printf("sum=%v,sub=%v\n", sum, sub)
	fmt.Printf("圆的面积是：%.2f\n", c.area())
}
