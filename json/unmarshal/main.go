package main

import (
	"encoding/json"
	"fmt"
)

/*
本章节学习如何将 json 反序列化为各种数据类型。
需要注意的是，反序列化时务必引用序列化时所使用的数据格式
包括：
1. 结构体类型
2. map 类型
3. slice 类型
*/

type Person struct {
	Name    string `json:"person_name"`
	Age     int    `json:"person_age"`
	Address string
}

// 1. 反序列化 struct
func unmarshalSturct() {
	// str 一般来自网络传输或者文件读取，本来是不需要用转义的，注意键使用了 tag
	str := "{\"person_name\":\"tom\",\"person_age\":10,\"Address\":\"北京东城\"}"
	// 声明空的结构体对象，用来承接反序列化后的结果
	var person Person
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Println("error = ", err)
	}
	fmt.Printf("结构体反序列化后的结果是:%v, person.Name=%v\n", person, person.Name)
}

// 2. 反序列化 map
func unmarshalMap() {
	str := "{\"Address\":\"北京朝阳\",\"Age\":10,\"Name\":\"momo\"}"
	// 定义一个 map 数据类型和序列化之前一样
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("error = ", err)
	}
	fmt.Printf("map反序列化后的结果是%v\n", a)

}

// 3. 反序列化 slice
func UnmarshalSlice() {
	str := "[{\"Address\":\"北京朝阳1\",\"Age\":11,\"Name\":\"sisi1\"}," +
		"{\"Address\":\"北京朝阳2\",\"Age\":12,\"Name\":\"sisi2\"}]" //字符串太长时，可以选择这样的方式拼接一下
	// 定义一个 slice 数据类型和序列化之前一样
	var b []map[string]interface{}
	err := json.Unmarshal([]byte(str), &b)
	if err != nil {
		fmt.Println("error = ", err)
	}
	fmt.Printf("slice反序列化后的结果是%v\n", b)
}

func main() {
	unmarshalSturct() // 结构体反序列化后的结果是:{tom 10 北京东城}, person.Name=tom
	unmarshalMap()    //map反序列化后的结果是map[Address:北京朝阳 Age:10 Name:momo]
	UnmarshalSlice()  //slice反序列化后的结果是[map[Address:北京朝阳1 Age:11 Name:sisi1] map[Address:北京朝阳2 Age:12 Name:sisi2]]
}
