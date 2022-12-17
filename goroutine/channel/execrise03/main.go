package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
练习：
1. 创建一个Person 结构体，有字段 Name Age Address
2. 利用 rand 方法配合随机创建 10 个 Person 对象，并放入到 channel 中
3. 遍历 channel ，将所有的 Person 对象信息打印到终端
*/

type Person struct {
	Name    string
	Age     int
	Address string
}

var ( // 定义一个名字池
	names = []string{
		"tom",
		"jack",
		"domi",
		"yuri",
		"pype",
		"musi",
	}

	// 定义一个地址池
	addresses = []string{
		"夏威夷",
		"巴黎",
		"柏林",
		"都柏林",
		"北京",
		"齐齐哈尔",
	}
)

// 定义一个函数，随机创建 Person 对象
func createPerson(names []string, addresses []string, c chan Person, exit chan int) {
	rand.Seed(time.Now().UnixNano()) // 随机Seed种子
	name := names[rand.Intn(len(names)-1)]
	address := addresses[rand.Intn(len(addresses)-1)]
	age := rand.Intn(100)
	person := Person{
		Name:    name,
		Age:     age,
		Address: address,
	}
	c <- person
	exit <- 1
}

func main() {
	resChan := make(chan Person, 10)
	exit := make(chan int, 10)
	for i := 1; i < 10; i++ {
		go createPerson(names, addresses, resChan, exit)
	}
	for i := 1; i < 10; i++ {
		<-exit
	}
	close(resChan)
	for v := range resChan {
		fmt.Println(v)
	}
}
