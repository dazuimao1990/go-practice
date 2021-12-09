package main

import "fmt"

// 练习1 编写一个函数 swap ,传入 a b int 后，可以交换二者的值

func swap(a *int, b *int) {
	t := *a // 中间变量传递法
	*a = *b
	*b = t
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Printf("swap 后，a = %v, b = %v\n", a, b)
}
