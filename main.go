package main

import (
	"fmt"
	"go_learning_note/basic"
	"go_learning_note/user"
)

func main() {
	fmt.Println("Main: Hello World!")
	s := user.Hello()
	fmt.Printf("user: %v\n", s)

	// 变量及数据类型
	fmt.Println("=======GoLang 基础之变量及数据类型=====")
	basic.TestDataType()

	// 数字类型
	fmt.Println("=======GoLang 基础之数字类型=====")
	basic.TestNumber()

	// 字符串类型
	fmt.Println("=======GoLang 基础之字符串类型=====")
	basic.TestString()
}
