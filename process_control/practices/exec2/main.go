package main

import "fmt"

func main() {
	// 综合练习4 根据输入的月份和年份，求出该月的天数
	// 闰年判断：判断条件为 1. 年份可以被 4 整除，但不可以被 100 整除。 或者 2. 年份可以被 400 整除
	var year int
	var isRun bool
	var month byte
	fmt.Println("输入年份:")
	fmt.Scanln(&year)
	fmt.Println("输入月份:")
	fmt.Scanln(&month)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%v 年是闰年\n", year)
		isRun = true
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Printf("%v 月有 31 天\n", month)
	case 2:
		if isRun {
			fmt.Printf("%v 月有 29 天\n", month)
		} else {
			fmt.Printf("%v 月有 28 天\n", month)
		}
	default:
		fmt.Printf("%v 月有 30 天\n", month)
	}

	// 综合练习5 根据输入的身高体重，来返回身体健康情况
	// 公式： （身高 - 108）* 2 = 体重 ，可以有10斤左右的浮动
	var hight int
	var wight int
	fmt.Println("请输入身高，单位cm")
	fmt.Scanln(&hight)
	fmt.Println("请输入体重，单位斤")
	fmt.Scanln(&wight)
	if wight-10 <= 2*(hight-108) && wight+10 >= 2*(hight-108) {
		fmt.Println("您的身材适中")
	} else {
		fmt.Println("您的身材过重或过轻")
	}

	// 综合练习6 判断此考试是什么等级
	// 90-100 优秀
	// 80-89  优良
	// 70-79  良好
	// 60-69  及格
	// 60以下  不及格

	var score float64
	fmt.Println("请输入分数")
	fmt.Scanln(&score)
	if score < 0 || score > 100 {
		fmt.Println("输入有误")
	} else {
		switch {
		case score >= 90:
			fmt.Println("优秀")
		case score >= 80:
			fmt.Println("优良")
		case score >= 70:
			fmt.Println("良好")
		case score >= 60:
			fmt.Println("及格")
		default:
			fmt.Println("不及格")

		}
	}
}
