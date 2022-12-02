package main

import "fmt"

func main() {
	// 顺序查找，就是在数组中遍历，比对元素和查找值是否相等
	// 案例：输入一个名字，判断是否在 [金毛狮王、白眉鹰王、紫山龙王、青翼蝠王] 中
	// 方案1:

	names := [4]string{"金毛狮王", "白眉鹰王", "紫山龙王", "青翼蝠王"}
	var name = ""
	fmt.Println("输入一个名字...")
	fmt.Scanln(&name)
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Printf("找到了%v，下标为%v \n", name, i)
			break // 找到了就跳出循环
		} else if i == (len(names) - 1) {
			fmt.Printf("没找到%v\n", name)
		}
	}

	// 方案2（推荐）
	index := -1
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到了%v，下标为%v \n", name, index)
	} else {
		fmt.Printf("没找到%v\n", name)
	}
}
