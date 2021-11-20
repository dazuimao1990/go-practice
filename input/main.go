package main

import "fmt"

var (
	name   string
	age    byte
	sal    float32
	isPass bool
)

func main() {

	// 要求：可以从终端输入姓名、年龄、薪水、是否通过考试
	// 输入时，引用变量的地址值，可以影响到变量本身
	// //方式 1 : 使用 fmt.Scanln() 方法
	// fmt.Println("输入姓名")
	// //运行到 fmt.Scanln(&name) 会让程序停住，等待输入
	// fmt.Scanln(&name)
	// fmt.Println("输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("输入薪水")
	// fmt.Scanln(&sal)
	// fmt.Println("输入是否通过考试")
	// fmt.Scanln(&isPass)
	// //输出保存的结果
	// fmt.Printf("考生姓名是 %v\n年龄 %v\n薪水 %v\n是否通过 %v\n", name, age, sal, isPass)

	//方式2 : 使用 fmt.Scanf() 方法
	fmt.Println("依次输入姓名、年龄、薪水、是否通过考试，空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("考生姓名是 %v\n年龄 %v\n薪水 %v\n是否通过 %v\n", name, age, sal, isPass)
}
