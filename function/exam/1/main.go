package main

import "fmt"

// 循环输入年份和月份
// 如果判断月份语句输入的不正确 ， 不在 1-12 之间，则输出月份错误的语句
func main() {
	for {
		var year int
		var month int
		fmt.Println("输入年份:")
		fmt.Scanln(&year)
		fmt.Println("输入月份:")
		fmt.Scanln(&month)
		if month < 1 || month > 12 {
			fmt.Println("输入的月份有误")
			continue
		}
		fmt.Printf("输入的是 %v 年 %v 月\n", year, month)
	}
}
