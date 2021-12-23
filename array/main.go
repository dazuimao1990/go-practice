package main

import "fmt"

// 数组，可以包含多个同数据类型元素的集合
// 数组的地址可以通过 &intArr 来获取
// 数组的第一个元素的地址，就是数组的首地址（起点地址）
// 数组的各个元素之间地址值的变化，取决于元素的数据类型所占字节数，如 int 类型的元素占 8 字节，则每个地址值之间增加 8
// 数组各元素，是有默认值的，取决于元素的数据类型，如 int 类型的默认值就是 0

func main() {

	// 四种初始化数组的方式
	var numArr1 [3]int = [3]int{1, 2, 3}
	var numArr2 = [3]int{5, 6, 7}
	var numArr3 = [...]int{8, 9, 10}
	var numArr4 = [...]int{1: 800, 0: 900, 2: 1000}        // 注意下标的顺序
	numArr5 := [...]string{1: "tom", 0: "jary", 2: "mary"} // 类型推导
	fmt.Printf("numArr1 的值是%v\n", numArr1)
	fmt.Printf("numArr2 的值是%v\n", numArr2)
	fmt.Printf("numArr3 的值是%v\n", numArr3)
	fmt.Printf("numArr4 的值是%v\n", numArr4)
	fmt.Printf("numArr5 的值是%v\n", numArr5)

	// 案例1 鸡场里有 6 只鸡，体重分别为 1 3 3.5 4 5.3 50 KG，求体重平均值
	// 使用 数组 的方式解决

	// 1. 定义一个数组，包含 6 个元素，数据的类型是 float64
	var hens [6]float64
	// 2. 为每个元素赋值
	hens[0] = 1
	hens[1] = 3
	hens[2] = 3.5
	hens[3] = 4
	hens[4] = 5.3
	hens[5] = 50
	// 3. 遍历数组，求和
	var totalWight float64
	for i := 0; i < len(hens); i++ {
		totalWight += hens[i]
	}
	// 4. 用总和除以数组长度，求平均
	avgWight := fmt.Sprintf("%.2f", totalWight/float64(len(hens))) // %.2f 用于格式化输出的小数点保留位数， 浮点类型变量和整数类型变量不能直接运算
	fmt.Printf("鸡们的总体重为 %v , 平均体重 %v\n", totalWight, avgWight)

	// 数组的遍历方式可以使用 for - range 方法

	var heroes = [3]string{"宋江", "林冲", "鲁智深"}
	for index, value := range heroes {
		fmt.Printf("下标 %v, 姓名 %v\n", index, value)
	}

}
