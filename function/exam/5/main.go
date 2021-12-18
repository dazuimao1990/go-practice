// 打印以下效果：
// 1. 加法
// 2. 减法
// 3. 乘法
// 4. 除法
// 0. 退出
// 请选择:
// 如果输入1 ，则返回 10 + 5 ；输入2 ，则返回 10 - 5 ；如输入 0 则退出程序
package main

import "fmt"

func inPut() int {
	var n int
	fmt.Println("1. 加法")
	fmt.Println("2. 减法")
	fmt.Println("3. 乘法")
	fmt.Println("4. 除法")
	fmt.Println("0. 退出")
	fmt.Println("请选择")
	fmt.Scan(&n) // 此处不需要参数校验，传入的值如果不是 int 类型，n 的值将保持原默认值 0
	return n
}

func main() {
	n := inPut()
	switch n {
	case 1:
		fmt.Println("10+5=", 10+5)
	case 2:
		fmt.Println("10-5=", 10-5)
	case 3:
		fmt.Println("10*5=", 10*5)
	case 4:
		fmt.Println("10/5=", 10/5)
	case 0:
		fmt.Println("退出程序")
	}
}
