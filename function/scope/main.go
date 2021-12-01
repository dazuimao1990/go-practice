package main

import "fmt"

// 函数的作用域：
// 函数中的变量，是局部变量，在函数外不会生效
// func scope1() {
// 	// 函数内部的局部变量
// 	//var n1 int = 10
// }

// 基本数据类型和数组都是值传递的，默认进行了值拷贝，而非引用。故此，在函数内部的修改，不会影响到原来的值。
func scope2(n1 int) {
	n1 = n1 + 10
	fmt.Println("scope2 n1 = ", n1) // 30
}

// 如果希望函数内的计算，可以影响到函数之外，可以通过传递变量地址（指针），函数以指针的方式操作变量，效果上看，类似引用。
// 这里传入的参数，是指针类型，
func scope3(n1 *int) {
	// n1 本身代表一个内存地址， *n1 代表取地址中的值
	*n1 = *n1 + 10
	fmt.Println("scope3 n1 = ", *n1) // 30
}

func main() {

	// 对 scope1 函数中定义的 n1 变量，无法进行直接引用
	// fmt.Println("n1 = ",n1)  // undeclared name: n1 局部变量无法引用

	// 使用 scope2 函数，传入 n1 的值，函数中自行计算输出，而 main 函数中的 n1 值不会受到影响。
	n1 := 20
	scope2(n1)
	fmt.Println("main n1 = ", n1) // 20

	// 使用指针变量，传递地址给函数，函数直接操作内存中的地址，更改其中保存的值，故而函数以外的值也变了
	num := 20
	scope3(&num)
	fmt.Println("main num = ", num) // 30
}
