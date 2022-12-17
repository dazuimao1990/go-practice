package main

import (
	"fmt"
	"log"
	"time"
)

/*
本章节学习在某个协程报错 panic 时，如何使用 defer + recover 来处理问题，避免主线程退出
1. 声明两个不同的协程
2. 在一个协程中故意引入一个错误
3. 捕获错误，让另一个协程和主线程正常运行
*/

func proc1() {
	for i := 0; i < 10; i++ {
		fmt.Println("proc1输出", i)
		time.Sleep(time.Second)
	}
}

func proc2() {
	// 使用匿名函数结合 defer +  recover 来实现错误的捕获与处理，这样就不会 Panic 了
	defer func() {
		if err := recover(); err != nil {
			log.Println("Proc2 出现错误", err)
		}
	}()
	var myMap map[int]int
	// 造一个经典错误，使用引用类型数据但是不 make, 如不实现处理，则会出现以下报错
	// panic: assignment to entry in nil map
	// myMap = make(map[int]int)
	myMap[0] = 1
	fmt.Println(myMap)
}

func main() {
	go proc1()
	go proc2()
	for i := 0; i < 10; i++ {
		fmt.Println("Main输出", i)
		time.Sleep(time.Second)
	}
}
