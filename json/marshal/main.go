package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
本章节学习如何将不同的数据类型进行 json 序列化并输出
包括：
1. 结构体类型
2. map 类型
3. slice 类型
4. 基本类型（对基本类型进行 json 序列化没有什么实际意义，但是也可以这么做）
*/

// 定义一个函数，利用 encoding/json 包完成对任意输入的 json 序列化
// json.Marshal 函数可以接受 interface{} 空接口作为输入，即任意类型都可以输入进去，这个用法要掌握
func JsonForEverything(src interface{}) {
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v 序列化后= %s \n", src, data) // data 的数据类型是 []byte ,如使用 %v 输出，则输出 ASICII 码，不如使用 %s 转化为 string 输出
}

// 定义一个结构体，供后续序列化使用
// 借助事先定义好的 tag ，可以在序列化时替换原有的 key ，即用 "person_name" 代替 “Name”，这是反射机制实现的
type Person struct {
	Name    string `json:"person_name"`
	Age     int    `json:"person_age"`
	Address string
}

func main() {
	// 1. 生成一个 Person 结构体对象
	person := Person{
		Name:    "tom",
		Age:     10,
		Address: "北京东城",
	}

	// 2. 声明一个 map 并赋值
	// map 使用之前一定要先 make 来开辟内存空间
	m := make(map[string]interface{})
	m["Name"] = "momo"
	m["Age"] = 10
	m["Address"] = "北京朝阳"

	// 3. 声明一个 slice 并赋值，其元素类型为上面的 map ，即 slice 中包含多个类似上述的 map
	var slice1 []map[string]interface{}

	for i := 1; i < 3; i++ {
		m_i := make(map[string]interface{}) // map 的声明在每次循环中要重建，这是由于 map 是引用类型，如果始终修改同一个 m_i 会导致 slice1 中的元素始终一样
		m_i["Name"] = "sisi" + fmt.Sprintf("%d", i)
		m_i["Age"] = 10 + i
		m_i["Address"] = "北京朝阳" + fmt.Sprintf("%d", i)
		slice1 = append(slice1, m_i) // append 代替了 make
	}

	// 4. 声明一个基本类型 float64
	var f float64 = 123.456

	// 输出 struct json
	// 注意值类型和引用类型的区别
	JsonForEverything(&person) // &{tom 10 北京东城} 序列化后= {"person_name":"tom","person_age":10,"Address":"北京东城"}
	JsonForEverything(m)       // map[Address:北京朝阳 Age:10 Name:momo] 序列化后= {"Address":"北京朝阳","Age":10,"Name":"momo"}
	JsonForEverything(slice1)  // [map[Address:北京朝阳1 Age:11 Name:sisi1] map[Address:北京朝阳2 Age:12 Name:sisi2]] 序列化后= [{"Address":"北京朝阳1","Age":11,"Name":"sisi1"},{"Address":"北京朝阳2","Age":12,"Name":"sisi2"}]
	JsonForEverything(f)       // 123.456 序列化后= 123.456
}
