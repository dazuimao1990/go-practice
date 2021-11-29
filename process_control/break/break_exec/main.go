package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 练习1 随机 100 以内的两个数求和，求出当和第一次大于20 的时候的当前数。

	// 思路分析
	// 先会求随机数，赋值给两个变量。
	// 死循环，做和
	// 判断和是否大于20 ，是则跳出循环，打印当前两个随机数
	var n1 int
	var n2 int
	rand.Seed(time.Now().Unix())
	for {
		n1 = rand.Intn(100)
		n2 = rand.Intn(100)
		if n1+n2 >= 20 {
			fmt.Printf("%d + %d = %d > 20\n", n1, n2, n1+n2)
			break
		}
	}

	// 练习2 实现登录认证，有三次机会，如果用户名为 张无忌，密码为 888 提示登录成功，否则提示还有几次机会

	// 思路分析
	// 大的结构是3次循环
	// 控制台输入接住用户名和密码
	// 判断正确与否
	// 正确跳出循环，否则输出剩余循环次数

	var userName string
	var passWord string
	for i := 3; i > 0; i-- {
		fmt.Println("输入用户名")
		fmt.Scanln(&userName)
		fmt.Println("输入密码")
		fmt.Scanln(&passWord)
		if userName == "张无忌" && passWord == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("验证错误，还有%d次机会。\n", i-1)
		}
	}
}
