package main

import "fmt"

/*
本章节学习如何管道 channel 的基本使用方法
1. 管道本质是一种队列类型的数据结构，与栈不同的地方是先进先出
2. 管道的使用要首先进行初始化，并且明确是哪种类型的管道，即管道中的元素的数据类型要提前声明
3. 管道是一种引用类型，使用前需要 make
4. 管道的容量需要在 make 时声明，而且容量值不会动态变化，这意味着如果管道的长度（元素的个数）超过容量，则会引起 deadlock 报错
5. 管道的元素不允许小于0，即不可以取光所有元素后再取，会 deadlock。这种问题在使用协程的时候不用担心，限于手动取值。
*/

func main() {
	// 声明一个管道 chan ，其元素的数据类型是 int
	myChan := make(chan int, 3)
	fmt.Printf("myChan 的值=%v myChan 的地址是 %p \n", myChan, &myChan) //myChan 的值=0xc0000b4000 myChan 的地址是 0xc0000aa018
	fmt.Printf("myChan 的长度是%v 容量是%v\n", len(myChan), cap(myChan)) //myChan 的长度是0 容量是3

	// 向管道内写入数据
	myChan <- 1 // <- 操作符用于向管道内写入，或者从管道向变量写入
	myChan <- 2
	num := 3
	myChan <- num // 由变量写入管道
	//myChan <- 4 // 不允许，写入多于容量，会死锁
	fmt.Printf("myChan 的长度是%v 容量是%v\n", len(myChan), cap(myChan)) //myChan 的长度是3 容量是3

	// 从管道向外取值
	n1 := <-myChan
	fmt.Println("n1=", n1)                                        // n1=1 先进先出，取出第一个元素
	fmt.Printf("myChan 的长度是%v 容量是%v\n", len(myChan), cap(myChan)) //myChan 的长度是2 容量是3

	// 关闭管道使用内置函数 close() ，被关闭的 chan 无法被写入，但是可以继续被读取
	close(myChan)
	//myChan <- 4 // 不允许写入已经被关闭的 chan
	n2 := <-myChan
	fmt.Println("n2=", n2)                                        // n2=2 被关闭的管道依然可以读取
	fmt.Printf("myChan 的长度是%v 容量是%v\n", len(myChan), cap(myChan)) //myChan 的长度是1 容量是3

	// 遍历管道
	myChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		myChan2 <- i * 2
	}
	fmt.Printf("myChan2 的长度是%v 容量是%v\n", len(myChan2), cap(myChan2)) //myChan 的长度是100 容量是100

	// 不宜使用普通的 for 循环，这是因为管道每读取一个，len 就 -1
	// for i:=0;i<len(myChan2);i++{
	// }

	// 应该使用 for range，并且在进行遍历前，应该关闭管道再遍历，否则到最后遍历循环会等待管道继续写入
	// 请求等待是造成 deadlock 的原因之一
	close(myChan2)
	for v := range myChan2 {
		fmt.Println("v=", v)
	} // v=2 - v=200
}
