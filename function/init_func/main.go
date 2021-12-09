package main

import "fmt"

// 每个源文件中都可以有一个 init 函数，它会在main函数之前被 go 调用，像钩子函数。
// 通常用来完成初始化工作

// init 函数细节：
// 1. 如果一个文件中同时存在全局变量定义、init 函数 、main 函数，则执行的顺序是 变量定义 > init 函数 > main 函数
// 2. 如果 main.go 引用了 utils.go 文件，则执行顺序的先后是 变量定义(utils.go) > init 函数(utils.go) > main.go {变量定义 > init 函数 > main 函数}
func init() {
	fmt.Println("init func...")
}

func main() {
	fmt.Println("main func...")
}
