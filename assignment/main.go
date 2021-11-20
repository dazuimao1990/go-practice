package main

import "fmt"

func main() {

	// 赋值运算符包括：
	// = 基本赋值方式
	// +=  a += 1 ==> a = a + 1
	// -=  a +- 1 ==> a = a - 1
	// *=  a *= 1 ==> a = a * 1
	// /=  a /= 1 ==> a = a / 1
	// %=  a %= 1 ==> a = a % 1
	var a int
	a = 9
	fmt.Println("a 赋值为", a)

	// 推导式的多重赋值
	a, b := 5, 6
	fmt.Printf("a 赋值为 %v , b 赋值为 %v\n", a, b)

	// 利用中间变量 t 交换 a b 的赋值
	// t := a
	// a = b
	// b = t
	// fmt.Printf("交换后，a 赋值为 %v , b 赋值为 %v\n", a, b)

	// 不使用中间变量 交换 a b 的赋值
	a, b = b, a
	fmt.Printf("交换后，a 赋值为 %v , b 赋值为 %v\n", a, b) //
	// 复合赋值方式操作
	a += 7 // a = a + 7 ==> 13
	fmt.Println("a 赋值为 ", a)
}
