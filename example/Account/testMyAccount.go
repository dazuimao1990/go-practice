package main

import (
	"fmt"
)

/*
使用 **面向过程** 的方式方法编程，实现一个记账程序，基本界面如下：
----------------------家庭记账软件------------------
                      1. 收支明细
					  2. 登记收入
					  3. 登记支出
					  4. 退出程序
请选择(1-4):

*/

func main() {
	// 定义字符串变量来接受输入的值
	key := ""

	// 定义变量获取余额
	balance := 10000.0 // 默认余额有一个达不溜
	// 定义变量获取入账
	money := 0.0
	// 定义变量获取说明
	note := ""
	// 定义收支列表
	details := "收支\t余额\t入账\t说明" // 后续的每一笔出入帐，向字符串追加即可
	// 定义死循环，switch 根据输入值进行判断
a: // 为以下死循环打一个 label
	for {
		fmt.Println("\n\n\n----------------------家庭记账软件------------------")
		fmt.Println("                      1. 收支明细")
		fmt.Println("                      2. 登记收入")
		fmt.Println("                      3. 登记支出")
		fmt.Println("                      4. 退出程序")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key) // 为 key 传值
		switch key {
		case "1":
			fmt.Println("---------收支明细--------")
			fmt.Println(details)
		case "2":
			fmt.Println("收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("收入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("支出金额:")
			fmt.Scanln(&money)
			balance -= money
			fmt.Println("支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			fmt.Println("退出程序...")
			break a // 中断标签
		default:
			fmt.Println("输入有误")
		}
	}

}
