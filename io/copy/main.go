package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 编写一个文件拷贝函数，接收源文件地址和目标文件地址
// 利用 io.Copy 包实现

func CopyFile(detFilePath string, srcFilePath string) (Written int64, err error) {
	// io.Copy 函数需要传入两个源文件和目标文件的文件句柄
	// 此处利用 bufio 包实现创建两个文件句柄
	srcFile, err := os.OpenFile(srcFilePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("Open Src file err : ", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	detFile, err := os.OpenFile(detFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Create Dest file err : ", err)
	}
	writer := bufio.NewWriter(detFile)
	return io.Copy(writer, reader)
}

func main() {
	// 调用函数，实现将 ../example/test.txt 拷贝为 ../example/test_copy.txt
	src := "../example/test.txt"
	dest := "../example/test_copy.txt"
	_, err := CopyFile(dest, src)
	if err != nil {
		log.Fatal("Cpoy file err: ", err)
	} else {
		fmt.Println("Copy success!")
	}
}
