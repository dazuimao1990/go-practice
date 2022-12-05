package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 引入带缓冲区的读取操作，读入文件中的所有内容
// bufio 包提供方法来创建一个带缓冲区（默认大小4096）的 *Reader 对象
// 可以用于读取文件

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

	// 创建 *Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 每读取到 \n 代表读到换行符，即 read line 操作
		if err == io.EOF {                  // io.EOF 代表读取到了文件的最后，这个东西会被作为错误处理也是我始料未及
			fmt.Println(str) // 此处的处理是因为发现如果文件中最后一行未换行，则不会打印出最后一行
			fmt.Println("文件读取完毕")
			break // 读取到最后，该退出退出
		}
		fmt.Print(str)
	}

}
