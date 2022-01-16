package main

import "fmt"

// 编写一个函数， fbn(n int), 要求可以接收一个 n int 能够将斐波那契数列放进一个切片中

// 1 声明一个函数，输入类型和返回类型确定好，使用 []uint64 的原因是，uint64 类型数据可以存放最大的值
func fbn(n int) []uint64 {
	// 使用 make 内置函数制作动态切片，容量 n 正好可以用作输入值的大小
	sliceFbn := make([]uint64, n)
	// 斐波那契数列开头两个值均为1， 作为递归函数的边界条件
	sliceFbn[0] = 1
	sliceFbn[1] = 1
	// 使用 for 循环和递归函数的形式得出指定长度的斐波那契数列
	for i := 2; i < n; i++ {
		sliceFbn[i] = sliceFbn[i-1] + sliceFbn[i-2]
	}
	// 返回切片
	return sliceFbn
}

func main() {

	fmt.Println(fbn(20))
}
