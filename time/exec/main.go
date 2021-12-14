package main

import (
	"fmt"
	"strconv"
	"time"
)

// 练习1 获取程序执行消耗的时间
func test() {
	str := ""
	for i := 1; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix()
	test()
	stop := time.Now().Unix()
	fmt.Printf("程序执行时间为%v秒\n", stop-start)
}
