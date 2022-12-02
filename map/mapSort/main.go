package main

import (
	"fmt"
	"sort"
)

func main() {
	// Map 在 golang 中是无序的，每次的输出都不太一样
	// 而且没有特定的方法对其进行排序
	// 如果需要依据 key 来对 map 排序，则先把 key 放入一个切片中，然后对切片排序，再用 key 输出对应的 value
	map1 := make(map[string]int)
	map1["aasd"] = 0
	map1["b"] = 10
	map1["zoo"] = 1
	map1["101"] = 100
	fmt.Println(map1) // 理论上此时的输出每次都不一样，然而在我的 go 1.17.3 中按顺序输出了···
	// 不重要，还是做一下排序
	// 声明一个 slice ，来放 map1 的 key
	var keys []string
	for k := range map1 {
		keys = append(keys, k)
	}
	// 使用 Sort 包中的方法为 keys 排序
	sort.Strings(keys)
	// 查看是否按照顺序输出
	fmt.Println(keys)
	// 遍历 keys ，来顺序输出 map1
	for _, v := range keys {
		fmt.Printf("map1[%v]=%v \n", v, map1[v])
	}

}
