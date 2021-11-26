package main

import "fmt"

// swtich 分支要点：
// 1. switch 的执行流程是，先执行表达式，得到值，然后从上到下依次和 case 表达式进行比较，如果相等，则执行对应的代码块，然后退出 switch
// 2. 如果 switch 的表达式没有和所有的 case 表达式匹配成功，则执行 default 语句块，然后退出 switch 控制
// 3. golang 的 case 后的表达式可以有多个，用 逗号 隔开
// 4. golang 的 case 对应语句块后面不需要写 break
// 5. case/switch 后面的表达式，可以是常量值（字面量）、变量、一个有返回值的函数
// 6. case 后面的表达式，和 switch 后的表达式，其值必须是同一个数据类型
// 7. case 后面如果是常量值，要求不可以重复（值在多个 case 后重复出现）
// 8. default 语句不是必须的，大不了啥都不执行退出
// 9. switch 后面可以啥都不加，当一个 if - else 使用，见练习2
// 10. switch 后面可以加赋值操作，用 ; 号结束，不推荐，见练习3
// 11. case 语句块后面可以加 fallthrough 关键字进行穿透，即这个语句块被匹配到后，执行完会继续执行下个 case 的语句块。
//     但不如直接把两个 case 合并，表达式逗号隔开，见练习4

// if 和 switch 的区别：
// 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用 switch ，简洁高效
// 其他情况，对区间判断和结果为 bool 值的判断。使用 if ，if的使用范围更广

func main() {

	// 练习1 输入字符 a b c d e f g ，分别对应输出 周一 周二...
	// 使用 switch 语句完成

	var key byte
	fmt.Println("请输入一个字符，a b c d e f g")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("周1")
	case 'b':
		fmt.Println("周2")
	case 'c':
		fmt.Println("周3")
	case 'd':
		fmt.Println("周4")
	case 'e':
		fmt.Println("周5")
	case 'f':
		fmt.Println("周6")
	default:
		fmt.Println("输入错了")

	}

	// 练习2 当 if - else 使用

	var score int = 91
	switch {
	case score >= 90:
		fmt.Println("good")
	case score >= 60:
		fmt.Println("pass")
	default:
		fmt.Println("nopass")
	}

	// 练习3 在 switch 后面加赋值语句

	switch grade := 99; {
	case grade >= 90:
		fmt.Println("good～")
	case grade >= 60:
		fmt.Println("pass～")
	default:
		fmt.Println("nopass～")
	}

	// 练习4 fallthrough 关键字进行穿透

	switch grades := 99; {
	case grades >= 90:
		fmt.Println("good~~")
		fallthrough // 会把后面的也执行
	case grades >= 60:
		fmt.Println("pass~~")
	default:
		fmt.Println("nopass~~")
	}

}
