package main

import (
	"fmt"
	"sync"
	"time"
)

/*
本章节学习如何利用全局互斥锁，解决多个协程同时向一块空间写入的问题
1. 利用 goroutine 求出1 - 20 的阶乘，并放到一个 map 里
2. 利用全局互斥锁，来解决写入 map 时的竞争问题
*/

var (
	myMap = make(map[int]int, 10)
	// sync: golang 中用于同步的包
	// Mutex: 互斥的意思，
	lock sync.Mutex
)

// 定义一个函数，用于计算整数 n 的阶乘 n!
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 将结果放入 map
	lock.Lock() // 锁住
	myMap[n] = res
	lock.Unlock() // 解锁

}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}

	time.Sleep(time.Second * 5) // 目前是比较鲁的方式来等待所有协程执行完毕，后续学习使用 channel 来同步
	for i, v := range myMap {
		fmt.Printf("%d!=%d\n", i, v)
	}
}
