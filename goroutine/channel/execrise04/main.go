package main

import (
	"fmt"
	"time"
)

/*
练习：
1. 求解 1 - 1000000 之间，哪些是素数
*/

/*
思路：
1. 启动协程，不断将数字写入一个 intChan
2. 定义函数，启动 100 个协程，从 intChan 中取出数字，计算判断是否为素数，如果是，则将数字写入一个结果管道 resChan
3. 利用 exit 管道，一旦100个协程都完成并向 exit 中写值，主线程就知道已经求解完成
4. 打印 resChan
*/

/*
一些思考：
最开始我使用了 slice 来承载结果，然而每次执行的结果都不一样，说明有问题。多个协程同时向slice 写入数据实测并不靠谱。
或许加入全局互斥锁可以解决问题，但全局互斥锁这种解决方案看起来会影响性能。
使用管道数据类型来承接结果，可以保障所有的协程写入时，是按照队列的形式压栈的，可以保障数据正确的写入
*/

func primeQury(n int) bool { // 最朴素的素数判断方法
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func cal(r chan int, c chan int, exit chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		if primeQury(n) {
			r <- n
		}
	}
	exit <- 1
}

func main() {
	intChan := make(chan int, 200000)
	resChan := make(chan int, 200000)
	resChan2 := make(chan int, 200000)
	exit := make(chan int, 100)

	go func() {
		for i := 1; i < 1000000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	start := time.Now().UnixMilli()
	for i := 1; i < 8; i++ {
		go cal(resChan, intChan, exit)
	}

	for i := 1; i < 8; i++ {
		<-exit
	}
	close(resChan)
	stop := time.Now().UnixMilli()
	fmt.Printf("协程统计素数共有%v个,计算耗时%v\n", len(resChan), stop-start)
	// 4 协程：协程统计素数共有78498个,计算耗时11757
	// 8 协程：协程统计素数共有78498个,计算耗时7591 (我的电脑有8个核心，超过8个协程就是并发+并行，少于或等于8个协程就是并行)
	// 100 协程：协程统计素数共有78498个,计算耗时7266 (经过测试，并发能够提供的优化，不如并行提供的多)

	// 对比传统单核方式计算时间
	start2 := time.Now().UnixMilli()
	for i := 1; i < 1000000; i++ {
		if primeQury(i) {
			resChan2 <- i
		}
	}
	close(resChan2)
	stop2 := time.Now().UnixMilli()
	fmt.Printf("单核统计素数共有%v个,计算耗时%v\n", len(resChan2), stop2-start2)
	// 单核统计素数共有78498个,计算耗时25135
}
