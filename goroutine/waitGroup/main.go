package main

import (
	"fmt"
	"sync"
	"time"
)

/*
本章节学习利用 sync.waitGroup 来实现多个协程之间的共同完成，替代 channel 章节中各个示例中的 exit 管道功能
1. 声明多个不同的协程
2. 声明一个 waitGroup 类型的变量
3. 启动每一个协程的时候，wg.Add(1) 来添加一个计数
4. 协程完成时调用 wg.Done 来减少一个计数
5. 主线程中调用 wg.Wait() 方法来等待所有的协程退出
6. 利用这种方式重写以下练习
*/

/*
练习：
1. 协程 writeData 向管道 intChan 内写入 50 个数据
2. 协程 readData 从管道 intChan 读取 50 个数据
3. 和主线程一起退出
*/

var (
	wg      sync.WaitGroup
	intChan chan int
)

func writeData(c chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		c <- i
	}
	close(c)
	fmt.Println("wariteData done:")
}

func readData(c chan int) {
	defer wg.Done()
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData read a value:", v)
		time.Sleep(time.Second / 5)
	}
}

func main() {
	intChan = make(chan int, 50)
	wg.Add(1)
	go writeData(intChan)
	wg.Add(1)
	go readData(intChan)

	// 等待 wg 归零
	wg.Wait()
	fmt.Println("所有的协程已经执行完毕，接下来是 Main 逻辑")
}
