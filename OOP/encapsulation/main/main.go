package main

import (
	"fmt"
	person "go-practice/OOP/encapsulation/model" // 引用包的时候，本地引用指向目录名字即可，目录内所有的包都会引用到
)

/* 一个封装示例：
1. 在另一个包中定义结构体 person 其中字段只有 Name 可以导出， age sel 字段不可导出
2. 定义 NewPerson 函数，用于为结构体创建对象实例
3. 为 person 中的 age sel 分别定义 SetXxx GetXxx 方法，为外部包引用操作该结构体实例的数据提供工厂模式的方法
*/

func main() {
	var p = person.NewPerson("tom")
	// 为年龄 薪水赋值
	p.SetAge(10)
	p.SetSel(3001.0)
	// 工厂模式调用得到年龄 薪水
	fmt.Printf("姓名：%v 年龄：%v 薪水：%v \n", p.Name, p.GetAge(), p.GetSel()) // 姓名：tom 年龄：10 薪水：3001

}
