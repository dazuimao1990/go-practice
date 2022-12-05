package main

import (
	"fmt"
	"log"
	"os"
	// "io/util" 在1.16版本后，该包被弃用，由 os.ReadFile 代替
)

func main() {
	// 使用此方法相当于一次性加载文件内容，故而只适合读写小文件
	// 不用显式的打开或者关闭文件
	content, err := os.ReadFile("../example/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n%s", content, content) // 改方式返回的数据类型为 []byte 直接输出的话，会转换成ASICII码，所以需要转化为 string 输出
}
