package main

import "fmt"

func main() {
	// 标准的 for 循环，i 为自增变量
	// for 循环基本语法：
	// for 循环变量初始化;循环条件;循环变量迭代 {
	// 	循环操作（语句）
	// }

	// for 循环的细节讨论
	// 循环条件的返回必须是个布尔值
	// 循环变量的作用域要注意，如果初始化在 for 语句中，那么作用域仅在 for 循环中，出了循环不可以引用。
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world", i)
	}

	// for 循环的第二种写法，只写循环条件，将循环变量的初始化和迭代写在别的地方
	// 循环变量的作用域要注意，如果初始化在 for 语句之外，那么作用域不再仅限于循环中，出了循环可以引用。
	j := 1
	for j <= 10 {
		fmt.Println("hello world~", j)
		j++
	}

	// for 循环的第三种写法，死循环，通常配合 break 使用
	// 下面是一个完全死循环，会导致资源耗尽的
	// for {
	// 	fmt.Println("hello world~~")
	// }
	// 通常配合 break 使用
	// 如下：
	k := 1
	for { // 等价于 for ; ; {
		if k <= 10 {
			fmt.Println("hello world~~~", k)
		} else {
			break
		}
		k++
	}

	// for 循环遍历字符串的方式

	// 第一种遍历字符串的方式，传统方式遍历，利用索引下标
	// var str string = "Hello,Goodrain!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i])
	// }
	// 这种方式有个缺陷，即str中包含中文字符时取值会有乱码。因为传统方式对字符串的遍历是按 **字节** 遍历，然而Golang默认使用的utf-8编码
	// 将每个汉字按照3字节存储，所以会出问题。
	// 解决方案有2种，第一种是使用 for-range 方式，可以直接输出中文汉字。
	// 第二种解决方式，是将字符串转换为切片，代码如下：
	var str string = "Hello,Goodrain!北京"
	// 重新定义新的变量，将字符串转换为切片
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	// 第二种遍历字符串的方式（推荐），for-range 方式
	// range str 可以返回两个值，分别对应索引和值，
	// for-range 按照 **字符** 遍历，是可以处理中文字符的
	// 由下例的输出，可以了解到，golang 中一个汉字占3个字节

	str = "abc_ok!~北京"
	for index, val := range str {
		fmt.Printf("index=%d , val=%c \n", index, val)
	}

	// Golang 语言中没有 while 和 do...while 循环，都是可以通过 for 循环构造出来一样的效果
	// 对于 while 循环而言，可以通过以下形式实现：
	var w int = 1
	for { // 死循环
		if w > 10 { // 循环跳出条件
			break // 跳出循环，跳出时， w = 11
		}
		fmt.Println("hello while", w)
		w++ // 循环变量迭代
	}
	fmt.Println("跳出循环时，w = ", w)

	// 对于 do...while 循环而言，可以通过以下形式实现：
	var do int = 1
	for {
		fmt.Println("hello do while", do) // 这种循环方式，与 while 之间的区别在于循环体至少执行一次
		do++
		if do > 10 {
			break
		}
	}
	fmt.Println("跳出循环时，do = ", do)
}
