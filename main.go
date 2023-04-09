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
	fmt.Println("\n=======GoLang 基础之变量及数据类型=====")
	basic.TestDataType()

	// 数字类型
	fmt.Println("\n=======GoLang 基础之数字类型=====")
	basic.TestNumber()

	// 字符串类型
	fmt.Println("\n=======GoLang 基础之字符串类型=====")
	basic.TestString()

	// 运算符
	fmt.Print("\n=======GoLang 基础之运算符=====\n")
	basic.TestOperator()

	// 流程控制
	fmt.Print("\n=======GoLang 基础之流程控制=====\n")
	basic.TestController()
}
