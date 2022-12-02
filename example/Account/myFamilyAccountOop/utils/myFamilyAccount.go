package familyAccount

import "fmt"

type familyAccount struct {
	// 定义字符串变量来接受输入的值
	key string
	// 定义变量获取余额
	balance float64
	// 定义变量获取入账
	money float64
	// 定义变量获取说明
	note string
	// 定义收支列表
	details string
}

func NewFamilyAccount() *familyAccount { // 工厂模式，完成对结构体*familyAccount的新建
	return &familyAccount{
		key:     "",
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t余额\t入账\t说明",
	}
}

// 实现方法，实现收支明细
func (f *familyAccount) showDetails() {
	fmt.Println("---------收支明细--------")
	fmt.Println(f.details)
}

// 实现方法，实现登记收入
func (f *familyAccount) income() {
	fmt.Println("收入金额:")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("收入说明:")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
}

// 实现方法，实现登记支出
func (f *familyAccount) pay() {
	fmt.Println("支出金额:")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("支出的太多，无法接受...")
		return
	} else {
		f.balance -= f.money
	}
	fmt.Println("支出说明:")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", f.balance, f.money, f.note)
}

// 实现方法，实现退出
func (f *familyAccount) exit() {
	fmt.Println("退出程序...")
}

// 实现方法，显示主页面
func (f *familyAccount) MainMeun() {
a: // 为以下死循环打一个 label
	for {
		fmt.Println("\n\n\n----------------------家庭记账软件------------------")
		fmt.Println("                      1. 收支明细")
		fmt.Println("                      2. 登记收入")
		fmt.Println("                      3. 登记支出")
		fmt.Println("                      4. 退出程序")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&f.key) // 为 key 传值
		switch f.key {
		case "1":
			f.showDetails() // 调用方法
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
			break a // 中断标签
		default:
			fmt.Println("输入有误")
		}
	}
}
