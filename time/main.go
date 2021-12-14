package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. 获取当前时间
	now := time.Now()
	// now 有个专门的数据类型，time.Time
	fmt.Printf("现在时间是 %v, now 的类型是%T\n", now, now)

	// 2. 通过 now 对应的方法，可以获取年月日、时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month())) // 月份的英文可以强转 int 显示
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3. 格式化时间
	// 第一种方式，呆呆的格式化打印
	fmt.Printf("年月日时分秒： %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("年月日时分秒： %d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v dateStr的类型是 %T\n", dateStr, dateStr)

	// 第二种方式，time.Format("2006-01-02 15:04:05")
	// 奇葩设定，记住一个确定的时间，想改成什么格式就改成什么格式
	// 1月2号，下午3点4分5秒，2006年
	fmt.Println(now.Format("2006.01.02 15/04/05"))
	fmt.Println(now.Format("2006.01.02"))
	fmt.Println(now.Format("15/04/05"))

	// 4. 时间常量
	// time.Hour      // 一个小时
	// time.Minute    // 60s
	// time.Second    // 1000ms
	// time.Millisecond // 1000微秒
	// time.Microsecond //1000 纳秒
	// time.Nanosecond
	// 需求，每隔 0.1 秒打印一个数字，到 100 退出
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100) // 能乘整数，不能除
		if i == 100 {
			break
		}

	}

	// 5. Unix 时间戳与 UninxNano 时间戳
	// 代表从 1970-01-01 00:00:00 到现在的秒数、纳秒数
	fmt.Printf("Unix 时间戳 %v， UnixNano 时间戳 %v\n", now.Unix(), now.UnixNano())
}
