package main

import (
	"fmt"
	"log"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 如果 os.Stat 不返回错误，则说明文件或文件夹存在
		return true, nil // 返回 true 以及空错误
	}
	if os.IsNotExist(err) { // os.IsNotExist 用于判断给定的错误是文件不存在这种错误，则说明文件或文件夹不存在
		return false, nil // 返回 false 以及空错误
	}
	return false, err // 如果上述正常的判断都漏过了，则说明发生未知错误，返回 false 以及真实的错误

}

func main() {
	var filePath string = "/tmp"
	res, err := PathExist(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if res {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}

}
