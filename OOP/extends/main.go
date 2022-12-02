package main

import "fmt"

/*
面向对象编程-继承
在 Golang 中体现为 A struct 可以嵌套有名 B 或匿名 C 的其他结构体（前者实现也叫组合）
一旦完成了继承，A struct 可以访问 B C 结构体内部的字段或方法
如果 B C 中有同名的字段或方法，而 A 中恰好没有定义该字段或方法，则调用时必须加名字
如果 A 中也有字段或方法，则采用就近原则，访问 A 的
A 同时继承了 B C ，称之为多重继承，尽量不要使用
*/

/*
示例：
类型：
小学生 Pupil
大学生 Graduate
都继承学生 Student
*/

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("Name = %v Age = %v Score = %v \n", s.Name, s.Age, s.Score)
}

type Pupil struct {
	Student // 匿名结构体嵌套
	Name    string
}

type Graduate struct {
	stu   Student // 有名结构体嵌套，后续调用时必须加名字
	Class string
}

func main() {
	var p Pupil
	p.Name = "Tom"          // 定义给了 Pupil.Name
	p.Age = 10              // 由于 Pupil 没有 Age 字段，所以定义给了 Pupil.Student.Age
	p.Student.Name = "Aryi" // 定义给了 Pupil.Student.Name
	p.ShowInfo()            // Name = Aryi Age = 18 Score = 0
	fmt.Println(p)          //{{Aryi 18 0} Tom}

	var g = &Graduate{ // 这种定义要注意格式，符合 field_1:value_1 , field_2:value_2 的总体格式， value_1 又嵌套为 stu 赋值一个 Student 类型的值
		stu: Student{
			Name:  "jack",
			Age:   18,
			Score: 100,
		},
		Class: "ad",
	}
	g.stu.Age = 18
	// g.Name = "jaja" // 错误用法，使用组合时，调用必须指定结构体的名字，正确用法如下
	g.stu.Name = "jaja"
	fmt.Println(*g) // {{jaja 18 100} ad}
}
