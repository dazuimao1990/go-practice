package main

import "fmt"

func main() {

	// 案例：从输入接收一个人的年龄，如果年龄大于18岁，那么输出“你的年龄大于18岁了，请为你的行为负责”
	var age byte
	fmt.Println("输入你的年龄:")
	fmt.Scanf("%d", &age)
	// golang 要求判断条件不写（）虽然不报错，但是不建议
	if age >= 18 {
		fmt.Println("你的年龄大于18岁了，请为你的行为负责")
	} else {
		fmt.Println("你的年纪不大，这次放过你了")
	}

	// Golang 中允许在分支判断语句中声明变量，以下是个例子
	// if age2 := 24; age2 >= 18 {
	// 	fmt.Println("你的年龄大于18岁了，请为你的行为负责")
	// }
}
