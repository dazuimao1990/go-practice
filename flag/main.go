package main

import (
	"flag"
	"fmt"
)

// 利用 flag 包解析命令行参数，使以下的命令行可以生效：
// ./main -u root --pwd pass -h=127.0.0.1 -P 3456

type config struct {
	user string
	pwd  string
	host string
	port int
}

func main() {

	var Config config
	flag.StringVar(&Config.user, "u", "", "用户名，默认为空")
	flag.StringVar(&Config.pwd, "p", "", "密码，默认为空")
	flag.StringVar(&Config.pwd, "pwd", "", "密码，默认为空") // 可以接收 -p abc || --p abc || --pwd abc || --pwd=abc
	flag.StringVar(&Config.host, "h", "localhost", "主机，默认localhost")
	flag.IntVar(&Config.port, "P", 3306, "端口，默认3306")

	// 必要步骤，调用解析 flag 参数，这一步必须在使用所有变量前
	flag.Parse()
	// 输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v\n",
		Config.user, Config.pwd, Config.host, Config.port) // user=root pwd=pass host=127.0.0.1 port=3456
}
