package main

import "fmt"

func main() {

	// 练习1 打印出 1～100 之间，所有 9 的倍数的整数的个数，以及它们的和
	// 思路分析
	// 定义两个变量，个数 counter 和 sum
	// 循环遍历 1～100，循环变量为 i
	// 对 9 取余数进行判断
	// 如果余数为 0 ，则 counter +1 ， sum 累计 sum + i
	var counter int = 0
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			counter++
			sum += i // 等价于sum = sum + i
		}
	}
	fmt.Printf("1-100 之间，共有 %d 个整数可以被 9 整除，它们的和是 %d\n", counter, sum)

	// 练习2 完成 6 的正整数和式分解，即输出 0 + 6 = 6 \n 1 + 5 = 6 \n 以此类推， 6 这个值可以变
	// 思路分析
	// 定义和变量 j
	// 定义一个循环变量 i，遍历 0 - j
	// 另一个变量等于 j - i
	// 构造输出
	var j int = 9
	for i := 0; i <= j; i++ {
		fmt.Printf("%d + %d = %d\n", i, (j - i), j)
	}
}
