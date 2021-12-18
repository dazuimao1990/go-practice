package main

import (
	"errors"
	"fmt"
)

func test() {
	// 故意制造一个错误的函数
	// 使用 defer + recover 进行错误处理
	// defer 后加入匿名函数
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}() // 结尾处的 () 对匿名函数发起调用，这个函数不需要传参，留空即可
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)

}

// 自定义错误处理
// 模拟一个函数，读取配置文件信息，如果配置文件的名字错误，则返回错误信息
func readConf(name string) (err error) {
	if name == "init.conf" {
		return nil
	} else {
		return errors.New("配置文件名不匹配")
	}
}

func confHandle(name string) {
	err := readConf(name)
	if err != nil {
		panic(err) // 抛出自定义的错误信息，并 终止 退出程序
	}

	fmt.Println("执行 confHanle 后面的代码。。。")
}

func main() {

	// 调用错误的函数，在不加异常处理时，会导致程序运行崩溃
	// 加入 defer + recover 异常处理后，即可打印错误信息的同时，继续执行后续的代码
	// test()
	// for {
	// 	fmt.Println("继续执行其他代码")
	// 	time.Sleep(time.Second)
	// }

	// 测试自定义错误处理
	confFileName := "init2.conf"

	confHandle(confFileName)
	fmt.Println("执行 main 后面的代码。。。")
}
