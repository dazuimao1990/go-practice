package main

import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func(int) int {
	var n int = 10           // 从这里开始，但是这个 n=10 只会初始化一次，对象的每次重复调用，匿名函数返回的 n 会供下次继续调用
	return func(x int) int { // 返回一个匿名函数，这个匿名函数引用了函数之外的 n，因此，n 和匿名函数组成一个闭包
		n = n + x // 经过这里
		return n  // 经过这里
	} // 一直到这里，就是一个闭包
	// 1. 闭包可以看作一个类，匿名函数就是操作（方法），n 理解为字段（属性）
}

// 案例1
// 1. 编写一个函数 makeSuffix(suffix string) 可以接收一个后缀名（比如 .jpg） 并返回一个闭包
// 2. 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀名，则加上，否则原样输出
// 3. 要求使用闭包实现
// 4. strings.HasSuffix ，这个函数可以判断某个字符串是否含有指定的后缀
func makeSuffix(suffix string) func(string) string {
	// 传入的 suffix 变量，和下面返回的匿名函数组成了闭包
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

func main() {

	// 闭包：
	// 闭包就是一个函数，和与其相关的引用环境（变量）组合的一个实体（整体）

	// 调用一次累加器，执行了函数 AddUpper() ，实际上返回了匿名函数给 f
	// f 就是闭包生成的一个对象了，所以以下每次调用，都是在 f 基础上再次操作（累加）
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	// 案例1 调用
	// 调用函数 makeSuffix 返回闭包
	// 使用闭包的好处：
	// 它保存了函数的状态，即使其属性发生的变化得以保留
	// 使用传统方式也可以实现这个案例，但是每次都要重新输入 jpg
	f2 := makeSuffix(".jpg")
	fmt.Println("处理后的文件名 = ", f2("guox"))
	fmt.Println("处理后的文件名 = ", f2("golang.jpg"))
	f2 = makeSuffix(".png") // 重新赋值，相当于重新调用，属性发生变化，后续的闭包调用实现了不同的效果
	fmt.Println("处理后的文件名 = ", f2("guox"))
	fmt.Println("处理后的文件名 = ", f2("golang.jpg"))
	fmt.Println("处理后的文件名 = ", f2("new.png"))
}
