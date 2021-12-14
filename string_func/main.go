package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 字符串是不可变的，以下所有的变化，如果以变量的方式传参，那么原变量的值是不会变的

	// 统计字符串长度用 len() 按 字节 返回
	// golang 中默认使用 utf-8 编码，一个ascii码（字母 数字）占一个字节，一个汉字占 3 个字节，一个 emoji 表情占 4 个字节
	var str string = "hello北🐛" // 5 + 3 + 4 = 12
	fmt.Println("str 的长度", len(str))

	str1 := "hello北京"
	// 字符串遍历，同时处理中文字符，需要将字符串转换为 rune 切片
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 ：%c\n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误：", err) // 错误的判断可以用来进行输入校验
	} else {
		fmt.Println("转换的结果是：", n)
	}

	// 整数转字符串
	str = strconv.Itoa(123456)
	fmt.Printf("type is %T, str=%v\n", str, str)

	// 字符串转 byte 切片
	var bytes = []byte("i am you")
	fmt.Printf(" bytes = %v\n", bytes)

	// []byte 转字符串 码值转字符并拼接
	str = string([]byte{97, 98, 99})
	fmt.Printf(" str = %v\n", str)

	// 10进制转 2 8 16 32 进制，返回字符串
	str = strconv.FormatInt(1929, 2)
	fmt.Printf("1929 的2进制是%v\n", str)
	str = strconv.FormatInt(1929, 32)
	fmt.Printf("1929 的32进制是%v\n", str)

	// 查找字符串中是否含有指定的子字符串，返回 true/false
	b := strings.Contains("a am guox ", "gu")
	fmt.Printf("b=%v\n", b)

	// 统计字符串中出现了几次指定的子串，返回 int
	num := strings.Count("cheese", "e")
	fmt.Printf("num=%v\n", num)

	// 字符串比较
	// == 区分大小写比较，返回 true/false
	// strings.EqualFold() 不区分大小写，返回 true/false
	b = strings.EqualFold("abc", "Abc") // 验证码
	fmt.Printf("b=%v\n", b)
	fmt.Println("大小写不一致时:", "abc" == "ABC")

	// 返回字符串子串，在整个字符串中第一次出现的 index 值，如果没有，则返回 -1
	index := strings.Index("NLP_aaa", "a") // 0 1 2 3 4
	fmt.Println("index = ", index)

	// 返回子串在字符串中最后一次出现的 index 值
	index = strings.LastIndex("go gohellolang", "go")
	fmt.Println("index = ", index)

	// 将指定的字符串中的子串 替换成为另外一个子串 strings.Replace("go go helo","go","java","n")
	// n 指定替换几个， n=-1时代表全部替换
	str = strings.Replace("gogogo helo", "go", "java", 2)
	fmt.Println("str = ", str)

	// 按照某个分割符字符，将字符串切分为数组
	strArr := strings.Split("helo,BEIj,北京", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
	}

	// 将字符串全部转化为大写/小写
	// strings.ToLower(str) strings.TpUpper(str)
	str = "Hello Golang"
	fmt.Println("转化小写：", strings.ToLower(str))
	fmt.Println("转化大写：", strings.ToUpper(str))

	// 去除字符串两边的空格
	fmt.Println("去除左右两边的空格", strings.TrimSpace("  spcal asdquh   "))

	// 去除字符串两边指定的字符串，下面的实例相当于去掉 "!" 和 " "
	fmt.Println("去掉两边指定的字符串", strings.Trim("! hello oollsss !", " !"))
	fmt.Println("去掉左边指定的字符串", strings.TrimLeft("! hello oollsss !", " !"))
	fmt.Println("去掉右边指定的字符串", strings.TrimRight("! hello oollsss !", " !"))

	// 判断字符串是否以指定字符开头/结尾
	fmt.Printf("字符串是否以指定字符串开头: %v\n", strings.HasPrefix("www.rainbond.com", "www"))
	fmt.Printf("字符串是否以指定字符串结尾: %v\n", strings.HasSuffix("www.rainbond.com", ".com"))

}
