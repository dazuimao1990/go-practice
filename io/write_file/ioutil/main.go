package main

import (
	"log"
	"os"
	// "io/util" 在1.16版本后，该包被弃用，由 os.ReadFile os.WriteFile 代替
)

func main() {
	// 将一个文件中的内容，读取后输入到另一个文件中去
	var fromFilePath string = "../../example/test.txt"
	var toFilePath string = "../../example/dest.txt"
	data, err := os.ReadFile(fromFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(toFilePath, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
