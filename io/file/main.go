package main

import (
	"fmt"
	"os"
)

func main() {

	// 打开文件使用 os.Open() 函数
	var filePath string = "../example/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开错误:", err)
	}

	// 文件打开后一定记得使用 file.Close() 关闭，这里可以结合 defer 关键字解决
	// defer 结合匿名函数
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("文件关闭失败:", err)
		} else {
			fmt.Println("文件已关闭")
		}
	}()

	// 文件变量的本质是指针，以下可以证明
	fmt.Printf("文件变量file=%v \n", file)

}
