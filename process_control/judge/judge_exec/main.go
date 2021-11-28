package main

import "fmt"

func main() {

	// 练习1 编写一段代码，声明两个 int32 变量并赋值，判断两数之和，如果结果大于 50 ，则输出“hello world”
	var a int32 = 256
	var b int32 = 144
	if a+b > 50 {
		fmt.Println("hello world")
	}

	// 练习2 编写一段代码，声明两个 float64 变量并赋值，如果第一个数大于10.0，且第二个数小于20.0，则输出两数之和
	var c float64 = 10.345
	var d float64 = 15.525
	if c > 10.0 && d < 20.0 {
		fmt.Println(c + d)
	}

	// 练习3 编写一段代码，声明两个 int32 变量并赋值，判断二者的和，如果可以同时被 3 5 整除，则打印提示信息
	var e int32 = 7
	var f int32 = 8
	if (e+f)%5 == 0 && (e+f)%3 == 0 {
		fmt.Printf("二者和是 %d,可以被3 5 同时整除\n", e+f)
	}

	// 练习4 编写一段代码，判断年份是否是闰年，判断条件为 1. 年份可以被 4 整除，但不可以被 100 整除。 或者 2. 年份可以被 400 整除
	var year int
	fmt.Println("输入年份")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("年份 %d 是闰年\n", year)
	} else {
		fmt.Printf("年份 %d 不是闰年\n", year)
	}
}
