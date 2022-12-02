package main

import "fmt"

/*
一个景区根据人的年龄不同，收取不同的票价。
年龄大于 18 收 20元，
其他情况免票。
编写结构体 Visitor 根据年龄收取门票并输出。
eg：
请输入姓名: 张飞
请输入年龄: 19
张飞的年龄为 19，收取门票 20 元。
*/

type Visitor struct {
	Name string
	Age  byte
}

func (v *Visitor) charge() (fee int) {
	if v.Age > 18 {
		return 20
	} else {
		return 0
	}
}

func main() {
	var v Visitor
	fmt.Println("请输入姓名:")
	fmt.Scanf("%s", &v.Name)
	fmt.Println("请输入年龄:")
	fmt.Scanf("%v", &v.Age)
	fmt.Printf("姓名:%s 年龄:%v 收费:%d元\n", v.Name, v.Age, v.charge())
}
