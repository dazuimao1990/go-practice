package main

import "fmt"

// go 中支持自定义数据类型，相当于给原有数据类型起了别名，但是 go 从语法上依然认为这是两个数据类型
type myInt int

// 函数类型也支持这样的定义
type myFunType func(int, int) int // myFunType 就代表了 func(int, int) int 这种函数数据类型

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar myFunType, n3 int, n4 int) int { // myFunType 就代替了 func(int, int) int
	return funvar(n3, n4) // 当我们需要在很多地方都定义这种函数类型时，一旦有修改，只需要改 type 定义即可
}

func main() {

	var n1 myInt = 40
	n2 := int(n1) // 即使 myInt 就是 int 类型，但是依然需要显式的进行数据类型转换
	fmt.Printf("n1 的数据类型是 %T ,值是%v。n2 的数据类型是 %T ,值是 %v\n", n1, n1, n2, n2)

	res := myFun(getSum, 10, 22)
	fmt.Println("res = ", res)
}
