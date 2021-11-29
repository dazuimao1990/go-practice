package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 案例1 随机生成 1 - 100 之间的整数，直到生成 99 ，记录生成多少次。
	// 思路分析，写一个死循环，当生成的值是 99 时，判断跳出循环

	// 小知识点，如何生成随机数
	// math/rand 包中，函数 rang.Intn(n) 可以生成 [0,n) 的伪随机整数
	// 为了生成一个彻底的随机数，我们需要为 rand 设置一个种子
	// time.Now().Unix() : 返回一个从 1970年1月1日 0时0分0秒 到现在的秒数，可以用来作为随机种子

	var count int = 0
	rand.Seed(time.Now().Unix()) // 随机种子放在了 for 的外面，只需要生成一次种子，真随机会一直进行下去
	for {
		n := rand.Intn(100) + 1 // rand.Intn(n) 生成的区间是 0 - 99，所以要加1
		count++
		fmt.Println("n=", n)
		if n == 99 {
			break // 跳出循环
		}
	}
	fmt.Println("执行次数:", count)

	// break 语句注意事项：
	// 1. break 如果在多层嵌套循环中，默认跳出离他最近的 for 循环
	// 2. break 如果在嵌套循环中，可以通过指定标签来跳出指定的循环，下面是个案例
label1:
	for i := 0; i < 4; i++ {
		//label2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break // 默认跳出离他最近的 for 循环
				break label1 // 由于在判断成立时，一次性跳出了最外边的循环，所以结果只输出 j=0\nj=1
			}
			fmt.Println("j=", j)
		}
	}
}
