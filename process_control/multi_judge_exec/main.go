package main

import (
	"fmt"
	"math"
)

// 首先要注意到，计算出的值大概率不是整数，所以给变量时要给浮点值进行运算
// 使用 float 系列的数据类型赋值时，记得值有小数点
var (
	a float64 = 2.0
	b float64 = 4.0
	c float64 = 2.0
)

func main() {
	// 练习 1
	fmt.Println("求方程 ax^2 + bx +c = 0 的根")
	// 当使用的数据类型比较容易晕的时候，类型推导真滴好用
	m := b*b - 4*a*c
	if m > 0 {
		// 计算时，一定要注意 () 的使用， 如果是 2 * a 的话，就错了
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("方程有两个实根，分别是 %v %v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("方程有一个实根，是 %v\n", x1)
	} else if m < 0 {
		fmt.Printf("方程无解")
	}

	// 练习2 女子选老公，三个条件 身高超过180 ，财产超过1000w ， 长得帅
	// 条件都满足： 立刻嫁
	// 条件至少一个满足： 嫁了吧，比上不足比下有余
	// 条件一个都不满足：滚

	var hight int
	var money float32
	var handsome bool
	fmt.Println("你多高？多少公分")
	fmt.Scanln(&hight)
	fmt.Println("你有多少钱？人民币多少钱")
	fmt.Scanln(&money)
	fmt.Println("你帅不？回到我 true or false")
	fmt.Scanln(&handsome)
	if hight >= 180 && money >= 10000000 && handsome {
		fmt.Println("立刻嫁")
	} else if hight >= 180 || money >= 10000000 || handsome {
		fmt.Println("嫁了吧，比上不足比下有余")
	} else {
		fmt.Println("滚")
	}

}
