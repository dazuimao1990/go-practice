package main

import (
	"fmt"
)

func main() {
	students := make(map[string]map[string]string)
	students["1"] = make(map[string]string) // 使用 map 必须 make
	students["1"]["name"] = "tom"
	students["1"]["sex"] = "male"
	students["1"]["address"] = "beijing"

	students["2"] = make(map[string]string)
	students["2"]["name"] = "mary"
	students["2"]["sex"] = "female"
	students["2"]["address"] = "shanghai"
	fmt.Println(students)
	fmt.Println(students["1"])
	fmt.Println(students["1"]["name"])

	// 一个比较复杂的 for-range 循环遍历
	for k1, v1 := range students {
		fmt.Println("studentNum=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v, v2=%v\n", k2, v2)
		}
	}

	// 使用 len() 函数统计 map 长度
	fmt.Println(len(students))
	fmt.Println(len(students["1"]))
}
