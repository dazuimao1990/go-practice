package main

import (
	"fmt"
)

/*
练习：
1. 协程 writeData 向管道 intChan 内写入 50 个数据
2. 协程 readData 从管道 intChan 读取 50 个数据
3. 和主线程一起退出
*/

// 向管道写入数据，并在写入完成后关闭管道
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

// 读取管道中的数据，并且向另一个管道中写入结果表示读取完毕，已通知主线程退出
func readData(intChan chan int, exitChan chan bool) {
	for { // 死循环读数据
		v, ok := <-intChan
		if !ok { // 如果读取失败，则认为读取完成
			break
		}
		fmt.Println("读取 v=", v)
	}
	exitChan <- true // 为退出管道传值，告知主线程退出
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 一种比较稳妥的方式，是死循环读管道，读到即退出
	// for {
	// 	_, ok := <-exitChan
	// 	if ok {
	// 		break
	// 	}
	// }

	<-exitChan // 管道是可以阻塞的，如果没有被关闭，那么可以在主线程中持续读取，知道得到值
}
