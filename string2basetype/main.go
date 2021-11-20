package main

import (
	"fmt"
	"strconv"
)

func main() {
	str2 := "123"
	str3 := "123.312"
	str4 := "true"
	var n1 int64
	var b bool
	var n2 float64

	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf(" n1 `s typr is %T and it`s value is %d\n", n1, n1)

	n2, _ = strconv.ParseFloat(str3, 64)
	//n3 := float32(n2)
	fmt.Printf(" n2 `s typr is %T and it`s value is %f\n", n2, n2)

	b, _ = strconv.ParseBool(str4)
	fmt.Printf(" b `s typr is %T and it`s value is %t\n", b, b)

}
