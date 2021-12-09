package main

import "fmt"

var (

	// 全局变量中定义匿名函数，可以在任意地方调用
	Func = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	// 匿名函数的三种用法

	// 1. 定义函数时直接调用，这种方式的匿名函数只能调用一次，是最常见的用法
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1 = ", res1)

	// 2. 匿名函数可以交给一个变量，这个变量就是函数类型，可以反复调用
	res2 := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println("res2 =", res2(20, 30))

	// 3. 匿名函数全局调用
	res3 := Func(2, 3)
	fmt.Println("res3 =", res3)
}
