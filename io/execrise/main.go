package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// 目标：设计一个程序，统计一个文件中，英文字符、数字、空格和其他字符的个数

// 思路：
// 1. 定义一个结构体，用来保存统计到的字符类型个数
// 2. 打开文件，创建一个文件句柄
// 3. 创建带缓冲的 reader
// 4. 循环读入文件的每一行
// 5. 循环遍历每一行的内容，并使用 switch 语句处理统计

type count struct {
	enCount    int
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func main() {
	filePath := "../example/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var counter count
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str { // string 是一个[]byte 所以在判断的时候，可以用 byte 来直接匹配
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 无条件落入下一个case中的执行体中去
			case v >= 'A' && v <= 'Z':
				counter.enCount++
			case v == ' ' || v == '\t':
				counter.spaceCount++
			case v >= '0' && v <= '9':
				counter.numCount++
			case unicode.Is(unicode.Han, v): // 利用 unicode 包的函数来判断是否为汉字字符
				counter.chCount++
			default:
				counter.otherCount++
			}
		}

	}
	fmt.Printf("英文字符的个数:%v,中文字符数量:%v,数字个数:%v,空格数量:%v,其他字符:%v.\n",
		counter.enCount, counter.chCount, counter.numCount, counter.spaceCount, counter.otherCount)
}
