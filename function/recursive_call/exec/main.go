package main

import "fmt"

// 练习1 利用递归函数，求出第 n 个斐波那契数。
// 1 1 2 3 5 8 13 ...
// 思路：
// 1. 当 n==1 || n==2 时，返回 1
// 2. 当 n>=2 时，返回 f(n-1) + f(n-2)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

// 练习2 求函数值，已知 f(1)=3 f(x)=2*f(x-1)+1 ，使用递归思想，求出 f(n) 的值
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

// 练习3 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个。以后每一天，猴子都
// 吃其中的一半，再多吃一个。 当第十天的时候，发现就只有一个桃子了，问最初有几个桃子
// 思路分析：
// 构造一个函数，自变量是天数 day ，因变量是桃子的数量
// f(day)=2 * (f(day+1)+1) 这个值，从第十天的值倒推第九天，更好归纳出规律
// 递归边界条件 f(10)=1

func monkyEatPeach(n int) int {
	if n == 10 {
		return 1
	} else if n >= 1 && n <= 10 {
		return 2 * (monkyEatPeach(n+1) + 1)
	} else {
		fmt.Println("输入的天数有误")
		return 0 // 表征返回错误的值
	}
}

func main() {

	var n int
	fmt.Println("输入正整数 n")
	fmt.Scanln(&n)
	fmt.Printf("第%v个斐波那契数是%v\n", n, fbn(n))
	fmt.Printf("f(%v) = %v\n", n, f(n))
	fmt.Printf("最开始共有桃子 %v\n", monkyEatPeach(1))

}
