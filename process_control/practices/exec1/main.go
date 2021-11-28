package main

import "fmt"

func main() {
	// 综合练习1 实现判断一个整数，属于什么范围，大于0；等于0；小于0
	var num int
	fmt.Println("请输入一个整数:")
	fmt.Scanf("%d", &num)
	if num > 0 {
		fmt.Printf("整数 %d 大于0\n", num)
	} else if num < 0 {
		fmt.Printf("整数 %d 小于0\n", num)
	} else {
		fmt.Printf("整数 %d 等于0\n", num)
	}

	// 综合练习2 判断整数是否为水仙花数，水仙花数指一个三位数，其各个位的立方和等于三位数本身，比如： 153 = 1*1*1 + 5*5*5 + 3*3*3
	// 分析思路：
	// 将三位数的百位、十位、各位求出并分别赋值给变量
	// 计算后做判断
	var num1 int
	fmt.Println("请输入一个三位整数:")
	fmt.Scanf("%d", &num1)
	hundreds := num1 / 100
	tens := (num1 - hundreds*100) / 10
	digits := num1 - hundreds*100 - tens*10
	if hundreds*hundreds*hundreds+tens*tens*tens+digits*digits*digits == num1 {
		fmt.Printf("%v hundreds %v tens %v digits, %v 是个水仙花数\n", hundreds, tens, digits, num1)
	} else {
		fmt.Printf("%v hundreds %v tens %v digits, %v 不是个水仙花数\n", hundreds, tens, digits, num1)
	}

	// 综合练习3 保存输入的用户名和密码，判断用户名是否为张三，密码是否为1234，如果是，则输出登录成功，否则提示登录失败
	var userName string
	var passWord string
	fmt.Println("请输入一个用户名:")
	fmt.Scanln(&userName)
	fmt.Println("请输入一个密码:")
	fmt.Scanln(&passWord)
	if userName == "张三" {
		if passWord == "1234" {
			fmt.Println("登录成功")
		} else {
			fmt.Println("错误的密码")
		}
	} else {
		fmt.Println("登录失败")
	}
}
