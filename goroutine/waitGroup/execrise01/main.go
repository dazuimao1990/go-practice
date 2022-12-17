package main

import (
	"fmt"
	"sync"
)

/*
练习：
1. 启动一个协程，将1-200 放入管道 numChan
2. 启动 8 个协程，从 numChan 中取出数字 n ，并计算 1+···+n，将结果放到 resChan 中
3. 8 个协程完成工作后，遍历 resChan 显示结果 ：res[1] = 1 .. res[10]=55 .. res[200]=?
4. 考虑 resChan 的数据类型用什么比较合适
*/

var wg sync.WaitGroup
var numChan chan int
var resChan chan map[int]int

func initpipe(c chan int) {
	defer wg.Done()
	for i := 1; i <= 200; i++ {
		c <- i
	}
	close(c)
}

func cal(c chan int, r chan map[int]int) {
	defer wg.Done()
	for {
		v, ok := <-c
		if !ok {
			break
		}
		res := 0
		for i := 0; i <= v; i++ {
			res += i
		}
		tmpMap := make(map[int]int)
		tmpMap[v] = res
		r <- tmpMap
	}
}

func main() {
	numChan = make(chan int, 20)
	resChan = make(chan map[int]int, 200)

	wg.Add(1)
	go initpipe(numChan)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go cal(numChan, resChan)
	}

	// 等待所有的协程完成工作
	wg.Wait()
	// 遍历前记得关闭管道
	close(resChan)
	for res := range resChan {
		for k, v := range res {
			fmt.Printf("res[%d]=[%d]\n", k, v)
		}
	}

}
