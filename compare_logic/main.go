package main

import "fmt"

// 测试函数，只要执行，就会打印，并返回 true
func test() bool {
	fmt.Println("触发 Test 函数")
	return true
}

func main() {
	// 关系运算符（比较运算符）包括 == != < > <= >= 返回的值永远是 true or false
	a, b := 2, 1
	fmt.Println("a = b ?", a == b)
	fmt.Println("a != b ?", a != b)
	fmt.Println("a < b ?", a < b)
	fmt.Println("a <= b ?", a <= b)
	fmt.Println("a > b ?", a > b)
	fmt.Println("a >= b ?", a >= b)
	flag := a < b
	fmt.Printf("flag is %v\n", flag)

	// 逻辑运算符：
	// && 代表逻辑与 左右两侧均为真时，结果为真；有不为真时，结果为假
	// || 代表逻辑或 左右两侧有真则结果为真；均为假时，结果为假
	// ! 代表取反
	A := 30
	if A > 30 && A < 40 {
		fmt.Println("30 < A < 40")
	}

	if A >= 30 && A <= 40 {
		fmt.Println("30 <= A <= 40")
	}

	if A > 30 || A < 40 {
		fmt.Println("30 < A or A < 40")
	}

	if A >= 30 || A <= 40 {
		fmt.Println("30 <= A or A <= 40")
	}

	if A > 30 {
		fmt.Println("30 < A")
	}

	if !(A > 30) {
		fmt.Println("A <= 30")
	}

	// 短路与
	// && 逻辑运算中，只要左侧为假，就不会继续执行后续的条件
	if !(A != 30 && test()) {
		fmt.Println("不触发 test函数，且触发短路与")
	}

	// 短路或
	// || 逻辑运算中，只要左侧为真，就不会继续执行后续的条件
	if A == 30 || test() {
		fmt.Println("不触发 test函数，且触发短路或")
	}

}
