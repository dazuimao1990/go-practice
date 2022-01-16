package main

import "fmt"

func main() {
	// 声明一个数组
	var intArr [5]int = [5]int{1, 22, 33, 44, 55}
	// 声明一个切片
	slice := intArr[1:3]
	slice1 := intArr[:3]
	slice2 := intArr[1:]
	slice3 := intArr[:]
	// 切片看作是数组的引用
	// 引用 intArr 数组起始下标为 1 终止下标为 3 的元素们，注意，包左不包右，不含 3
	fmt.Println("数组 intArr 为", intArr)
	fmt.Println("切片 slice 的元素为", slice)      // [22 33]
	fmt.Println("切片 slice 的长度为", len(slice)) // 2
	fmt.Println("切片 slice 的容量为", cap(slice)) // 4 容量的值取自底层数组 intArr 的末位元素下标
	fmt.Println("切片可以不声明start，默认是0 ", slice1)
	fmt.Println("切片可以不声明end，默认是len(intArr) ", slice2)
	fmt.Println("切片可以不声明start和end，默认是0和len(intArr) ", slice3)

	for i, v := range slice3 {
		fmt.Printf("第%v个元素是%v\n", i, v)
	}

	// 切片可以继续切片
	slice4 := slice[0:2]
	fmt.Println("切片可以继续切片", slice4)

	// 切片是引用类型，slice4 变化了，基于原数组而来的所有切片都会变化，数组也会变化
	slice[0] = 10086
	fmt.Println("数组 intArr 为", intArr)
	fmt.Println("切片 slice 的元素为", slice)      // [22 33]
	fmt.Println("切片 slice 的长度为", len(slice)) // 2
	fmt.Println("切片 slice 的容量为", cap(slice)) // 4 容量的值取自底层数组 intArr 的末位元素下标
	fmt.Println("切片可以不声明start，默认是0 ", slice1)
	fmt.Println("切片可以不声明end，默认是len(intArr) ", slice2)
	fmt.Println("切片可以不声明start和end，默认是0和len(intArr) ", slice3)

	// 通过 append 方法扩展切片，实际上是重新分配了一个数组
	slice = append(slice, 101, 102, 103)
	fmt.Println("扩展后的切片为", slice)
	// 通过 append 向切片追加切片（下面的格式是固定的， slice...）
	slice = append(slice, slice...)
	fmt.Println("扩展后的切片为", slice)

	// 使用 copy 拷贝切片
	var sliceSrc []int = []int{1, 2, 3, 4, 5}
	sliceDst := make([]int, 10)
	fmt.Println("sliceSrc", sliceSrc)
	fmt.Println("sliceDst", sliceDst)
	// 反人类的设计，拷贝的 target在前面
	// 两个参数必须是切片类型
	copy(sliceDst, sliceSrc)
	// 拷贝关系的两个切片是完全独立的两个切片，源的变化不影响目标切片
	// 如果源切片长度大于目标切片长度，拷贝也不会报错，目标切片不会扩容，完成已有长度元素的复制后结束
	sliceSrc[0] = 10086
	fmt.Println("sliceSrc", sliceSrc)
	fmt.Println("sliceDst", sliceDst)

	// 切片与字符串的处理
	// 字符串底层是 byte 数组，所以可以作为切片处理
	var str string = "hello@goodrain"
	// 可以利用一个切片来接住最后几个字符
	slice5 := str[6:]
	fmt.Println("slice of str =", slice5)

	// 字符串是不可变类型，所以无法通过 str[0]='z' 的方式改变字符串中的元素
	// 如果需要改变字符串中的元素，需要先转为 []byte 或者 []rune 类型的切片，然后再转回字符串
	slice6 := []byte(str)
	slice6[0] = 'z'
	str = string(slice6)
	fmt.Println(str)
	// 然而 []byte 只能处理英文和数字
	// 汉字的处理需要使用 []rune
	slice7 := []rune(str)
	slice7[0] = '好'
	str = string(slice7)
	fmt.Println(str)
}
