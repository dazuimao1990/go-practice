package main

import "fmt"

func main() {

	// goto label 代表无条件跳转到带有指定 label 的语句去继续执行，中间的代码会被忽略
	// goto 一般和条件判断一起使用
	// goto 不被提倡，会使代码变得混乱，不易调试和理解

	var n int = 20
	fmt.Println("OK1")
	if n >= 20 {
		goto label // 代码一旦运行到这里，会直接跳转到 label 处继续执行
	}
	fmt.Println("OK2")
	fmt.Println("OK3")
label:
	fmt.Println("OK4")
	fmt.Println("OK5")

	// return 代表跳出当前的函数，即不再执行后续的代码，如同终止一个函数
	// return 一旦在 main 函数中，可以理解为退出程序
	for i := 1; i <= 10; i++ {
		if i == 3 {
			return
		}
		fmt.Println("循环运行次数", i)
	}
}
