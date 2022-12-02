package main

import (
	"fmt"
)

func main() {
	// 声明一个 map
	// map[keyType]valueType
	// map 使用之前必须 make 下
	// make 创建一个 map，且分配容量 10
	var a map[string]string = make(map[string]string, 10)

	a["no1"] = "songjiang"
	a["no2"] = "wuyong"
	a["no1"] = "wusong"
	a["no3"] = "wusong"

	// 类型推导赋值方法

	heros := make(map[string]string)

	heros["hero1"] = "guox"

	fmt.Println(a)
	fmt.Println(heros)

	// map 的若干特性
	// 1. map 使用前必须 make
	// 2. map 的 key 不可以重复，重复赋值则会使用最后一组
	// 3. map 的 value 可以重复
	// 4. map 内的元素是无序的

}
