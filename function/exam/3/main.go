// 编写一个函数，输出 100 以内的所有素数（仅可以被 1 和自己整除的数字），每行显示 5 个，并求和
package main

import "fmt"

// 判断一个数是否是素数
// 判断所有的结果取余不为0，不好判断
// 但是判断发现有一个因子不是素数好判断，因为仅需要一个反例
// 用 bool 值来反转我要的结果
func isPrime(n int) bool {
	// 2 是最小的素数，1 不是
	if n == 1 {
		return false
	}
	// 每次运算，循环2到n-1，如果出现可以被整除的则返回假，非素数
	for i := 2; i < n; i++ {
		// 发现有一个因子
		if n%i == 0 {
			return false
		}
	}
	// 上述条件筛选过后，剩下的就是素数
	return true
}

func main() {

	// 循环 100 次，对每个数进行运算

	// 每输出 5 次，则求和打印，并换行
	var sum int = 0
	var count int = 0
	for n := 2; n <= 100; n++ {
		if isPrime(n) {
			sum += n
			fmt.Printf("%v\t", n)
			if count++; count%5 == 0 {
				fmt.Println("五个素数和", sum)
				sum = 0
			}
		}
	}
}
