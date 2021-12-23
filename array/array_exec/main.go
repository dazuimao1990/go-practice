package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec01() {
	var arr01 [26]byte
	// 向 26 个 byte 元素的数组中存入 A-Z
	for i := 0; i < 26; i++ {
		arr01[i] = 'A' + byte(i) // 需要将 i 转为 byte 类型，才可以和 'A' 进行运算
	}

	fmt.Printf("26元素的数组是%c\n", arr01)
	// for index, value := range arr01 {
	// 	fmt.Printf("第 %v 个元素是 %c \n", index+1, value)
	// }
}

func exec02() {
	// 求出一个数组的最大值，并给出下标
	var arr02 [5]int = [5]int{1, 3, 5, 777, 87}
	var max int
	var index int
	for i := 0; i < len(arr02); i++ {
		if arr02[i] > max {
			max = arr02[i]
			index = i
		}
	}
	fmt.Printf("第 %v 个元素 %v 最大\n", index, max)
}

func exec03() {
	// 求出一组数组的和 和 平均值，使用 for-range
	var arr03 [5]int = [5]int{1, 3, 5, 777, 87}
	var sum int
	for _, value := range arr03 {
		sum += value
	}
	fmt.Printf("数组元素的和为 %v，平均值为 %v\n", sum, float64(sum)/float64(len(arr03))) // 平均值想保留小数，那么分数线上下都得是 float
}

func exec04() {
	// 随机生成 5 元素的数组，并在打印的时候，反转输出结果
	// 1. 生成随机种子
	rand.Seed(time.Now().UnixNano())
	// 2. 声明数组，并随机赋值
	var arr04 [5]int
	var length = len(arr04)
	for i := 0; i < length; i++ {
		arr04[i] = rand.Intn(100)
	}
	fmt.Println("随机生成的数组是:", arr04)
	// 3. 反转数组，交换数组头尾元素，共交换 length / 2 次 (5/2=2 6/3=3思考下)
	temp := 0
	for i := 0; i < length/2; i++ {
		temp = arr04[length-1-i]     // 将尾元素的值保存给临时变量
		arr04[length-1-i] = arr04[i] // 将尾元素赋值为头元素
		arr04[i] = temp              // 将临时变量赋值给头元素，完成交换
	}
	fmt.Println("反转后的数组是:", arr04)
}

func main() {
	exec01()
	exec02()
	exec03()
	exec04()

}
