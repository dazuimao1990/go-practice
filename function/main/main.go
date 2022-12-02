package main

import (
	"fmt"

	"go-practice/function/utils" // 调用包的路径，到文件夹，而不是包名
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

func myFunc(funcVar func(int, int) int, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

// 一个可以同时返回多个值的函数
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// 上面这个函数，可以支持提前将返回定义，即函数返回值可以命名
func getSumAndSub2(n1 int, n2 int) (sum int, sub int) { // 这个函数在被调用的时候，不需要写 return 的列表了
	sum = n1 + n2 // 这里不再是 := ,这是因为 sum 这个变量的类型已经在返回值列表中定义了，只需要重新赋值即可
	sub = n1 + n2
	return
}

// 函数支持可变参数，用 args... 来代表多个参数,args 是约定俗成的，可以改名字
// 如果一个函数具有可变参数，那么这个可变参数，要放在形参列表的最后一个
func sumMany(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i] // others 本质是个切片，others[i]代表第 i+1 个元素
	}
	return sum // 返回最终的结果
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

	// 调用有返回值命名的函数
	res4, _ := getSumAndSub2(19, 30)
	fmt.Printf("res4=%v\n", res4)

	// 函数也是一种数据类型，可以被赋给某个变量，则该变量就是一个函数类型的变量
	a := getSum
	fmt.Printf("a的类型是%T，getSum 的类型是%T\n", a, getSum)
	fmt.Println(a(10, 20)) // 等价于 getSum(10,20)

	// 函数既然是一种数据类型，那么也可以作为形参传递给另一个函数，并且调用，参考上面的函数 myFunc
	res5 := myFunc(getSum, 5, 6)
	fmt.Println(res5)

	// 调用有可变参数的函数
	res6 := sumMany(10, 1, 3, 4, 6, 7, 8)
	fmt.Println(res6)

}
