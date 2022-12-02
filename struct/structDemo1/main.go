package main

import (
	"fmt"
)

// 定义一个结构体 Cat，具有猫类的一系列 属性/字段

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	// 结构体（Cat）是自定义的数据类型，代表了一类事物
	// 结构体变量（实例）（cat1）是具体的，实际的，代表一个具体变量
	// 利用结构体声明一个 实例/对象 cat1 ，并赋给它对应的属性

	var cat1 Cat      // var a int //  Golang 的面向对象是天然的，
	fmt.Println(cat1) // 此时对象为空值，因为还没有赋值属性
	cat1.Name = "honey"
	cat1.Age = 3
	cat1.Color = "white"
	cat1.Hobby = "eat fish"
	fmt.Println("Cat1的详细信息：")
	fmt.Printf("cat1 Name: %v\n", cat1.Name)
	fmt.Printf("cat1 Age: %v\n", cat1.Age)
	fmt.Printf("cat1 Hobby: %v\n", cat1.Hobby)

	// 重新定义属性，重新赋值即可
	cat1.Name = "tom"
	fmt.Println("Cat1的详细信息：")
	fmt.Printf("cat1 Name: %v\n", cat1.Name)
	fmt.Printf("cat1 Age: %v\n", cat1.Age)
	fmt.Printf("cat1 Hobby: %v\n", cat1.Hobby)
}
