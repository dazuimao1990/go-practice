package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
本章节学习如何 go 并发编程
包括：
1. goroutine
2. channel
当前内容引出 goroutine :
1. 实现在主线程中打印10次 hello world
2. 同时在协程中同时打印10次 hello golang
*/

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello golang " + fmt.Sprintf("%d", i))
		time.Sleep(time.Second)
	}

}

func main() {
	// 查询逻辑cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Printf("我的 Mac 有 %v 个逻辑CPU\n", cpuNum)
	// 设定 golang 可用的 cpu 数量，下面的示例是预留一个逻辑cpu
	// golang 1.8 后默认让程序运行在多个 cpu 上，所以不再需要设置了
	runtime.GOMAXPROCS(cpuNum - 1)
	go test() // 一键开启协程，可以理解为将 test() 放到一边执行去了，代码继续向下执行下面的循环
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world " + fmt.Sprintf("%d", i))
		time.Sleep(time.Second)
	}
}
