package main

import "fmt"

func main() {

	// 练习1 continue 实现打印 1 - 100 之间的奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是", i)
	}

	// 练习2 从键盘输入若干个整数，最后统计正数个数和负数个数，要求使用for break continue
	var positive int
	var negative int
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positive++
			continue // 此处的 continue 意味着一旦判断到为正数，就不必进入后续的代码，直接跳出当前循环即可
		}
		negative++
	}
	fmt.Printf("有%v个正数，%v个负数\n ", positive, negative)

	// 练习3 某人有 100000 元，每经过一个路口都需要缴费
	// 当现金大于 50000 时，一次交 5%
	// 当现金小于等于 50000 时，一次交 1000
	// 求一共可以过几个路口
	var money float64 = 100000
	var crossNum int = 0
	for {
		if money < 1000 {
			break
		}
		if money > 50000 {
			crossNum++
			money = money - money*0.05
			fmt.Println("还剩下", money)
			continue
		}
		crossNum++
		money = money - 1000
		fmt.Println("还剩下", money)
	}
	fmt.Println("通过的路口数量", crossNum)
}
