package main

import (
	"fmt"
)

// defer 延迟执行，会在所属的函数执行完毕后执行的语句
// 最佳实践：
// 在打开或生成资源（文件句柄或数据库连接）后，加入 defer 语句，最后用来关闭资源
// 写入 defer 后，可以继续操作资源，因为 defer 关闭的操作，会在函数执行完毕后才执行
// 伪代码：
// func handlefile(){
// 	file = openfile(文件名)  //  开启文件句柄
// 	defer file.close()      // 压栈关闭句柄
// 	file.appand()           // 文件句柄操作
// }
func sum(n1 int, n2 int) int {
	// defer 使用注意事项：
	// 1. defer 会将对应的语句压入一个 defer 栈中，暂时不执行
	// 2. 当函数执行完毕时，会从栈中依据先入后出的规则，从 defer 栈中取出语句执行
	defer fmt.Println(" ok1 n1=", n1) // 第三个执行
	defer fmt.Println("ok2 n2=", n2)  // 第二个执行
	// 3. defer 入栈时会将相关的值拷贝压入栈中，在当前例子中，指代 n1 n2 ，在 defer 出栈执行时，其值还是压栈时的值，不受下面变动影响
	n1++
	n2++
	sum := n1 + n2
	fmt.Println(" ok3 res=", sum) // 第一个执行
	return sum
}

func main() {

	res := sum(10, 20)
	fmt.Println("fannel res=", res) // 最后执行
}
