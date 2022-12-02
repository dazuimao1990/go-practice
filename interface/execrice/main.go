package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
目标： 借助 sort.Sort 接口，将 Student 的切片，按照分数从大到小排列

https://pkg.go.dev/sort#Sort

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

// 1 定义结构体 Student
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 2 定义 Student 的切片类型
type StuSlice []Student

// 3 为 StuSlice 实现排序所需的 sort.Sort 接口
func (stu StuSlice) Len() int {
	return len(stu)
}

func (stu StuSlice) Less(i, j int) bool {
	// 按照分数进行排序
	return stu[i].Score > stu[j].Score
}

func (stu StuSlice) Swap(i, j int) {
	// 经典交换顺序写法
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	// 定义生成10个学生的结构

	var students StuSlice
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu-%d", i),
			Age:   rand.Intn(20),
			Score: rand.Float64() * 100.0,
		}
		// 将Student追加到切片
		students = append(students, stu)
	}

	// 遍历输出未排序的结果
	fmt.Println("=================== 未排序 ==================")
	for _, v := range students {
		fmt.Println(v)
	}
	// 开始排序，并输出
	sort.Sort(students)
	fmt.Println("=================== 排序后 ==================")
	for _, v := range students {
		fmt.Println(v)
	}

}
