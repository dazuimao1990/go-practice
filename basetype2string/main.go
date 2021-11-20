package main

import "fmt"

func main() {
	var num1 int = 876
	var num2 float32 = 123.333
	var b bool = true
	var myChar byte = 'z'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf(" str `s type is %T , and it`s value is %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf(" str `s type is %T , and it`s value is %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf(" str `s type is %T , and it`s value is %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf(" str `s type is %T , and it`s value is %q\n", str, str)
}
