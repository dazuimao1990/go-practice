package main

import (
	"fmt"
	"time"
)

/*
本章节学习在不关闭管道的情况下，如何利用 select 在多个管道之间协同工作，
1. 声明两个不同的管道
2. 利用 select 语句来同时接收管道的内容
3. 管道不会被关闭，但是也不会 deadlock
*/

func main() {
	// 声明俩不同的 chan
	intChan := make(chan int, 10)
	strChan := make(chan string, 5)

	// 为管道传值
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 为管道传值
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 死循环从管道中接收值
	for {
		select { // 在两个未关闭的管道中接收值
		case v := <-intChan:
			fmt.Println("从intChan中接收到了", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Println("从strChan中接收到了", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("不玩了，结束吧")
			return // 此处意味着直接退出 main 函数
		}
	}
}
