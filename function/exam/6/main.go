// 输出 a-z Z-A  ascii
package main

import "fmt"

func iToAsc(n int) {
	fmt.Printf("%c", n)
}

func main() {
	// 打印 a-z / 97-122
	for i := 97; i <= 122; i++ {
		iToAsc(i)
	}
	// 打印 Z-A / 90-65
	for i := 90; i >= 65; i-- {
		iToAsc(i)
	}
	fmt.Println()

}
