package main

import "fmt"

func main() {

	// 案例1 continue 的条件一旦成立，会跳出当前一次循环，继续下次循环，而不是跳出整个循环。
	// continue 支持基于标签跳出到指定的循环
label1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				//continue // 循环输出 j=0\nj=1\j=3 三次
				continue label1 // 跳出到最外层循环，所以 j=3 不会被输出了。 输出 j=0\nj=1 三次
			}
			fmt.Println("j=", j)
		}
	}
}
