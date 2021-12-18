package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 编写一个函数
// 随机生成 1--100 之间的整数
// 有十次机会猜
// 如果第一次猜中，则输出 你可真是个天才
// 如果 2--3 次猜中，则输出 还不错哦
// 如果 4--9 次猜中，则输出 一般般
// 如果 第10次猜中，则输出 你可算猜中了
// 一次都没有猜中，则输出，有点笨哦

func randNum(max int) int {
	rand.Seed(time.Now().Unix())
	var res int = rand.Intn(max) + 1
	return res
}

func enterNumVerify(n, max int) (err error) {
	if n >= 1 && n <= max {
		return nil
	} else {
		return fmt.Errorf("输入的数字应该位于 1-%v 之间，浪费1次机会", max)
	}
}

func enterNum(max int) (guress int) {
	var guess int
	fmt.Printf("输入 1 - %v 猜一下:\n", max)
	fmt.Scan(&guess)
	err := enterNumVerify(guess, max)
	if err != nil {
		fmt.Println(err)
	}
	return guess
}

func main() {
	max := 100
	guessChanceNum := 10
	res := randNum(max)

	guessed := false

	for i := 1; i < guessChanceNum; i++ {
		guess := enterNum(max) // 输入数字并校验，校验失败则退出
		if guess == res {
			guessed = true
			switch i {
			case 1:
				fmt.Println("你可真是个天才")
			case 2, 3:
				fmt.Println("还不错哦")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("一般般")
			case 10:
				fmt.Println("你可算猜中了")
			}
			break
		} else if guess < res {
			fmt.Printf("猜小了，还有%v次机会\n", guessChanceNum-i)
		} else if guess > res {
			fmt.Printf("猜大了，还有%v次机会\n", guessChanceNum-i)
		}
	}
	if !guessed {
		fmt.Println("有点笨哦，正确的值是", res)
	}

}
