package main

import "fmt"

func main() {

	// 嵌套循环知识：
	// 将一个循环放进另一个循环体内，就形成了一个嵌套循环。在外面 for 的称为外层循环，内部的 for 是内层循环
	// 嵌套循环不要超过 3 层
	// 实质上，嵌套循环就是把内层循环当作外层循环的循环体，当只有内存循环的循环条件为 false 时，才会跳出内层循环，结束当次外层循环，开始下次外层循环
	// 外层循环次数为 m ， 内层循环次数为 n，则内层循环最终要执行 m * n 次

	// 案例1 统计 3 个班的同学的成绩，每个班有 5 名同学。求出各个班级的平均分，和所有班级的平均分。学生成绩要求键盘输入。
	// 思路分析
	// （1）先易后难
	// （2）先死后活

	// 第一步
	// 1. 统计一个班级的 5 名同学的分数，求出该班的平均分
	// 2. 声明一个 sum 变量，用来计算总分
	// 3. 键盘输入

	// 第二步
	// 1. 统计 3 个班，每班 5 个同学的总分，平均分
	// 2. 声明变量 j 用来循环班级
	// 3. 声明变量存放所有班的总成绩

	// 第三步
	// 1. 把代码做活
	// 2. 声明两个变量，来代替班级数量，和学生数量

	// 追加需求，要求统计每个班的及格人数，以及总的及格人数
	// 声明变量 passCounter 用来保存各班及格人数
	// 声明变量 totalPassCounter 用来保存所有及格人数

	var classNum int = 2
	var stuNum int = 3
	var totalSum float64 = 0.0
	var totalPassCounter int = 0
	var passScore float64 = 60.0
	for j := 1; j <= classNum; j++ {
		var score float64 = 0.0
		var sum float64 = 0.0
		var passCounter int = 0
		for i := 1; i <= stuNum; i++ {
			fmt.Printf("请输入第%v班，第%v名同学的成绩：\n", j, i)
			fmt.Scanln(&score)
			sum += score
			// 判断及格人数以及及格分数线
			if score >= passScore {
				passCounter++
			}
		}
		fmt.Printf("第%v班的班级平均分%v; 共有%v名同学及格\n", j, sum/float64(stuNum), passCounter)
		totalSum += sum
		totalPassCounter += passCounter
	}
	fmt.Printf("所有班级总成绩是%v，总平均分%v,共有%v名同学及格\n", totalSum, totalSum/float64(stuNum*classNum), totalPassCounter)
}
