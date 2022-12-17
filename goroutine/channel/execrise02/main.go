package main

import "fmt"

/*
练习：
1. 启动一个协程，将1-200 放入管道 numChan
2. 启动 8 个协程，从 numChan 中取出数字 n ，并计算 1+···+n，将结果放到 resChan 中
3. 8 个协程完成工作后，遍历 resChan 显示结果 ：res[1] = 1 .. res[10]=55 .. res[200]=?
4. 考虑 resChan 的数据类型用什么比较合适
*/

/*
思路:
1. 生成 numChan 管道，放入200个数
2. 生成 resChan 管道，管道数据类型为 map
3. 定义函数，不停取出 numChan 元素，并计算取出值的 1+···+n 并将结果生成一个 map 并放入 resChan
4. 定义标志管道，在所有的计算完成后，向标志管道传递特征值
5. 主线程得到标志管道后，开始遍历 resChan
*/

// 生成目标管道 numChan
func initChan(numChan chan int) {
	for i := 1; i <= 200; i++ {
		numChan <- i
	}
	close(numChan)
}

// 定义函数，根据取到的数计算和，并将结果以 map 形式传递给 resChan，取无可取之时，为标志位 exit 管道 +1
func sum(numChan chan int, resChan chan map[int]int, exit chan int) {
	for { // 死循环不停取出
		v, ok := <-numChan
		if !ok {
			break // 取不到数据了，就退出
		}
		resNum := 0
		for i := 1; i <= v; i++ {
			resNum += i
		}
		tmpMap := make(map[int]int)
		tmpMap[v] = resNum
		resChan <- tmpMap
	}
	exit <- 1
}

// 主线程
func main() {
	numChan := make(chan int, 200)
	resChan := make(chan map[int]int, 200)
	exit := make(chan int, 8)

	// 调用协程，生成 numChan
	go initChan(numChan)

	// 启动 8 个协程，来计算 1+..+n
	for i := 1; i <= 8; i++ {
		go sum(numChan, resChan, exit)
	}

	// 主线程取标志管道结果，能取到 8 个值，说明所有的 8 个协程都已经完成任务退出，可以继续运行
	for i := 1; i <= 8; i++ {
		<-exit
	}

	// 遍历结果管道时，需要先关闭
	close(resChan)
	for res := range resChan {
		for k, v := range res {
			fmt.Printf("res[%d]=[%d]\n", k, v)
		}
	}

}
