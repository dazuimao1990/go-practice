package main

import "fmt"

// 冒泡排序
// 有个数组 [24,69,80,57,13] , 将其进行从小到大的排序
// 第 1 轮排序，找到数组中最大的值，放在最右边
// 从第 0，1 个元素开始比较，如果 0 > 1 则交换位置
// 持续到第 3，4 个元素，就找出了最大的值
// 第 1 次比较： [24,69,80,57,13]
// 第 2 次比较： [24,69,80,57,13]
// 第 3 次比较： [24,69,57,80,13]
// 第 4 次比较： [24,69,57,13,80]

// 第 2 轮排序，找到数组中第二大的值，放在倒数第二个位置
// 第 1 次比较： [24,69,57,13,80]
// 第 2 次比较： [24,57,69,13,80]
// 第 3 次比较： [24,57,13,69,80]

// 第 3 轮排序，找到数组中第三大的值，放在倒数第三个位置
// 第 1 次比较： [24,57,13,69,80]
// 第 2 次比较： [24,13,57,69,80]

// 第 4 轮排序，找到数组中第二大的值，放在倒数第四个位置
// 第 1 次比较： [13,24,57,69,80]

// 总结规律：
// 1. 进行 len(arr)-1 轮比较
// 2. 每轮比较，进行 len(arr)-1 次比较
// 3. 每次比较，需要按照大小，决定是否交换两个元素的位置

func BubbleSort(arr *[5]int) {
	// 传入一个数组，看看原来啥样
	fmt.Println("原数组:", *arr) // 数组是值类型，需要使用地址传入，在函数内部通过 *arr 来取值
	// 编写外层循环，来定义多轮比较
	// 定义临时变量，用于后续的交换操作
	for i := 0; i < len(arr)-1; i++ {
		temp := 0
		for j := 0; j < len(arr)-1-i; j++ { // len(arr)-1-i => 4,3,2,1
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后的数组:", *arr)
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
}
