package main

import "fmt"

func main() {
	var i int = 100
	// &i 取值为 i 变量在内存中的地址
	fmt.Printf("i 的值是 %v ,i 的地址是 %v\n", i, &i)
	// *int 表示声明的变量 ptr 的类型是一个面向 int 的指针
	// 赋值 &i 意味着把 ptr 赋值为 i 的内存地址
	// ptr 变量本身也有一个内存地址
	// *加内存地址，意味着取出内存地址中保存的值
	var ptr *int = &i
	fmt.Printf("ptr 的值是 %v , ptr 的地址是 %v \n", ptr, &ptr)
	fmt.Printf("ptr 代表的地址中，保存真正的值是 %v\n", *ptr)

	var j string = "asdasdasd"
	// &i 取值为 i 变量在内存中的地址
	fmt.Printf("j 的值是 %v ,j的地址是 %v\n", j, &j)
	// *int 表示声明的变量 ptr 的类型是一个面向 int 的指针
	// 赋值 &i 意味着把 ptr 赋值为 i 的内存地址
	// ptr 变量本身也有一个内存地址
	// *加内存地址，意味着取出内存地址中保存的值
	var ptrj *string = &j
	fmt.Printf("ptrj 的值是 %v , ptrj 的地址是 %v \n", ptrj, &ptrj)
	fmt.Printf("ptrj 代表的地址中，保存真正的值是 %v\n", *ptrj)

}
