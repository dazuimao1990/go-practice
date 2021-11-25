package main

import "fmt"

func main() {

	// 键盘输入学生的分数
	// 分数100，奖励 bmw 一台
	// 分数 (80,99]  奖励iphonePlus
	// 分数 [60,80]  奖励自行车
	// 其他分数，啥也没有

	var score int
	fmt.Println("输入分数：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("奖励 bmw 一台")
	} else if 80 < score && score <= 99 {
		fmt.Println("奖励iphonePlus")
	} else if 60 <= score && score <= 80 {
		fmt.Println("奖励自行车")
	} else {
		fmt.Println("奖励个粑粑")
	}
}
