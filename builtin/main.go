package main

import "fmt"

func main() {

	num1 := 100
	fmt.Printf("num1 的类型是 %T  num1 的值是 %v num1 的地址是%v\n", num1, num1, &num1)

	// new() 可以分配一个地址，用来保存类型，生成的值具有该类型的默认值
	num2 := new(int)

	// num2 的类型是 *int 指针
	// num2 的值是 地址 ，保存一个地址，这个地址中保存了新创建的int的值，默认值是0
	// num2 的地址是 地址，这个地址，是 num2 变量的地址
	fmt.Printf("num2 的类型是 %T  num2 的值是 %v num2 的地址是%v num2 地址指向的值是%v\n", num2, num2, &num2, *num2)

	*num2 = 200 // 给指针地址对应的值重新赋值
	fmt.Printf("num2 的类型是 %T  num2 的值是 %v num2 的地址是%v num2 地址指向的值是%v\n", num2, num2, &num2, *num2)
}
