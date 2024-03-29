package main

import (
	"fmt"
)

var (
	// i 是个正常的十进制整数
	i int = 396
	// j 是个八进制数，以 0 开头的就默认是 8 进制数
	// 八进制逢8进1(0 - 7 没有 8 ，8 是 010)，所以十进制的 9 ，相当于8进制的 010(8) + 1 ==> 011
	j int = 02456
	// k 是个十六进制数，以 0x 或者 0X 开头表示
	// 十六进制逢 16 进 1,(0 - 9 A(10) - F(15) , 16 用 0x10 表示)，所以十进制的 17，相当于十六进制的 0x10(16) + 1 ==> 0x11
	k int = 0xA45
)

func main() {

	// 如何将其他进制整数用二进制输出
	fmt.Printf("%b %b %b\n", i, j, k)
	// 使用十进制输出
	fmt.Printf("%d %d %d\n", i, j, k)

}
