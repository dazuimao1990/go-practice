package main

import "fmt"

func main() {

	// map 的增删改查
	// map["key"] = value 如果 key 不存在就是新增，如果存在就是修改

	cities := make(map[string]string)
	// key 不存在时是新增
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	fmt.Println(cities)

	// key 存在时为修改
	cities["no3"] = "天津"
	fmt.Println(cities)

	// map 中的查找
	val, ok := cities["no1"]
	fmt.Println(ok) // 发现找到时，ok被赋值为 true， 反之 false
	if ok {
		fmt.Printf("存在key%v\n", val)
	} else {
		fmt.Printf("不存在key no1\n")
	}
	// 删除 map 中的元素使用 delete 内置函数，delete(map,key)
	delete(cities, "no1")
	fmt.Println(cities)
	// 没有 map 或者 key 不会导致报错，只是不操作
	delete(cities, "no4")
	fmt.Println(cities)

	// 如果希望一次性清理 map 中的所有元素，有两种方式：
	// 1. 遍历所有的 key ，然后逐个删除
	for keys := range cities {
		fmt.Println(keys)
		delete(cities, keys)
	}
	fmt.Println(cities)

	// 2. 重新 make 一个同名的 map ，替换掉原来的，会被 gc 回收，简单粗暴有效
	// cities := make(map[string]string)
}
