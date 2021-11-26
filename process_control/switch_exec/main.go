package main

import "fmt"

func main() {

	// 练习1 使用 switch 将 a b c d e 转为大写，其他的输出 other
	var char byte
	fmt.Println("输入一个字符")
	fmt.Scanf("%c", &char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("Other")
	}

	// 练习2 对学生成绩大于等于 60 的输出及格，小于 60 的输出不及格，输入的成绩不可以大于 100
	var score int8
	fmt.Println("输入你的成绩")
	fmt.Scanln(&score)
	switch {
	case score >= 60 && score <= 100:
		fmt.Println("pass")
	case score >= 0 && score < 60:
		fmt.Println("nopass")
	default:
		fmt.Println("错误的输入，输入0 - 100")
	}

}
