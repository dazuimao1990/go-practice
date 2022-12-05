package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// 创建一个新文件，并写入 5 句 “hello world”
	// 定义文件路径
	var filePath string = "../../example/bufio_write.txt"
	// 打开文件，更一般的方式
	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644) // 只写模式，清空原内容
	// file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644) // 只写模式，在文件末尾追加
	// file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_APPEND, 0644) // 读写模式，并追加
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644) // 只写模式，没有文件则创建
	if err != nil {
		log.Fatal("open file err=", err) // 该操作会在触发错误时直接退出并打印日志
	}
	// 及时关闭文件
	defer file.Close()
	// 创建一个写文件的对象
	writer := bufio.NewWriter(file)
	str := "hello world\r\n" // \r\n 是为了应对有的编辑器换行用 /r 而不是 /n
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// bufio 带缓冲，所以文件写入后实际上是先在内存里的，需要进行以下操作来刷盘
	writer.Flush()

}
