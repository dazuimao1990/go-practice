package main

import (
	"fmt"
	familyAccount "go-practice/example/Account/myFamilyAccountOop/utils"
)

func main() {
	fmt.Println("这次是面向对象实现")
	familyAccount.NewFamilyAccount().MainMeun()
}
