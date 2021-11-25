package main

import "fmt"

var (
	second float32
	sex    string
)

func main() {

	// 练习1 输入百米跑成绩，少于 8s 的提示进入决赛，否则提示淘汰
	// 并提示输入性别，来决定进入男子/女子组决赛
	fmt.Println("输入你的百米成绩，秒:")
	fmt.Scanln(&second)
	if second <= 8 {
		fmt.Println("恭喜进入决赛组，请声明你的性别，man 或者 woman:")
		fmt.Scanln(&sex)
		if sex == "man" {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}
	} else {
		fmt.Println("很遗憾，你被淘汰了")
	}

	// 练习2 根据输入打印票价，
	// 4-6 月为旺季：
	// 成人（18-60） : 60
	// 老人 (> 60）: 半价
	// 儿童 (< 18）: 1/3
	// 其余月份为淡季
	// 成人 40
	// 其他 20
	var month byte
	var age byte
	fmt.Println("输入月份:")
	fmt.Scanln(&month)
	fmt.Println("输入你的年龄:")
	fmt.Scanln(&age)
	if month >= 4 && month <= 6 {
		if age < 18 {
			fmt.Println("请付款 20 元")
		} else if age <= 60 {
			fmt.Println("请付款 60 元")
		} else {
			fmt.Println("请付款 30 元")
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("请付款 40 元")
		} else {
			fmt.Println("请付款 20 元")
		}
	}
}
