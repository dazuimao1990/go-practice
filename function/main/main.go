package main

import (
	"fmt"

	"example.com/gogogo/function/utils" // 调用包的路径，到文件夹，而不是包名
)

// 函数是为了实现某个目标的代码集合
//
// func 函数名 (形参列表) (返回值列表){
// 	  执行语句
// 	  return 返回值列表
// }
//
// 形参列表： 表示函数的输入
// 函数中的执行语句：为了实现某一个目标的代码块
// 函数可以有返回值，也可以没有

// 函数调用的机制：
// 1. 在调用一个函数的时候，会在内存中分配一个新的空间（栈区）编译器会通过自身机制，让这个栈区和别的栈区空间分隔开来
// 2. 在每个空间中，数据空间是独立的，不会混淆在一起 （变量的作用域）
// 3. 当一个函数被调用（执行）完毕时，程序会销毁这个函数对应的栈区空间

// 一个仅返回一个值的函数
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ", sum)
	// 当函数有 return 语句时，就是将结果返回给调用者
	// 哪里调用返哪里
	return sum
}

// 一个可以同时返回多个值的函数
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	var n1 float64 = 10.0
	var n2 float64 = 5.0
	var operator byte = '/'

	// 调用引用其他包中的函数
	result := utils.Cal(n1, n2, operator) // 函数调用, 包名.函数名
	fmt.Println("result = ", result)

	n1 = 20.0
	n2 = 15.0
	operator = '+'
	// 调用引用其他包中的函数
	result = utils.Cal(n1, n2, operator) // 重复调用
	fmt.Println("result = ", result)

	// 调用当前包中的函数
	sum := getSum(10, 33) // sum 变量接收了函数的 return 值，这里等价于 sum := 30
	fmt.Println("main sum = ", sum)

	// 调用一个可以返回多个值的函数
	res1, res2 := getSumAndSub(19, 20)
	fmt.Printf("res1=%v,res2=%v\n", res1, res2)

	// 调用一个可以返回多个值的函数时，可以使用 _ 作为占位符，来忽略其中某个返回值
	_, res3 := getSumAndSub(19, 20)
	fmt.Printf("res3=%v\n", res3)

}
