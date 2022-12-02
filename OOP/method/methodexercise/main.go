package main

import "fmt"

/*
1 编写一个结构体 MethidUtils ，编程一个方法，不需要任何参数，在方法中打印一个 10 * 8 的矩形，在main函数中调用
*/
type MethidUtils struct {
	Row    int
	Column int
}

// 无传参，无返回示例
func (m MethidUtils) Output() {
	for i := 1; i <= m.Row; i++ {
		for j := 1; j <= m.Column; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 有传参，无返回示例
func (m MethidUtils) Area(len int, width int) {
	fmt.Println("矩形面积为:", len*width)
}

// 有传参，有返回示例
func (m MethidUtils) Area2(len float64, width float64) (str float64) {
	return len * width
}

func main() {
	me := MethidUtils{
		Row:    8,
		Column: 10,
	}
	me.Output()
	me.Area(3, 5)
	res := me.Area2(4, 8)
	fmt.Printf("第二个面积为%v\n", res)

}
