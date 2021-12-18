// 编写一个函数，三天打鱼两天晒网，判断是打鱼还是晒网，如果从 1990年1月1日起开始执行 “三天打鱼两天晒网”，如何判断以后的某一天中，是打鱼还是晒网
package main

import (
	"fmt"
	"time"
)

func subDays(t1, t2 time.Time) (days int) {
	days = int(t1.Sub(t2).Hours() / 24) // 计算两个时间的差值，以小时格式输出，故而除以 24 得到相差天数，并转换为 int 类型
	return days
}

func main() {

	var someDate string // 输入日期时，以字符串的格式输入更合理
	fmt.Println("输入一个日期，格式如：1990-09-27")
	fmt.Scan(&someDate)
	someDay, _ := time.Parse("2006-01-02", someDate) // 将字符串解析为时间类型，首个参数是时间格式，那个神奇的时间
	endDay, _ := time.Parse("2006-01-02", "1990-01-01")
	subDays := subDays(someDay, endDay)
	if subDays%5 < 3 { // 余 5 的结果，实际上是 0 1 2 3 4 ，所以取前 3 个使用 < 3
		fmt.Printf("%v 日距离1990-01-01 %v 天，应该打鱼\n", someDate, subDays)
	} else {
		fmt.Printf("%v 日距离1990-01-01 %v 天，应该晒网\n", someDate, subDays)
	}

}
