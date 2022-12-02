package person

import "fmt"

type person struct {
	Name string
	age  int // 小写开头的字段无法被外部包直接访问，封装住数据
	sel  float64
}

func NewPerson(name string) *person { // 返回结构体对象的指针
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) { // 方法实现过程中添加数据校验逻辑，通过时赋值，不通过时输出报错
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄输入错误。。。")
	}
}

func (p *person) GetAge() int { // 简单返回查询到的数据
	return p.age
}

func (p *person) SetSel(sel float64) { // 同 SetAge 方法
	if sel > 3000.0 && sel < 15000.0 {
		p.sel = sel
	} else {
		fmt.Println("薪水输入错误。。。")
	}
}

func (p *person) GetSel() float64 {
	return p.sel
}
