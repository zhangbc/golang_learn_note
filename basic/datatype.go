package basic

import (
	"fmt"
)

// 测试函数
func f() {

}

// 获取姓名 & 年龄
func GetNameAndAge() (string, int) {
	return "zhangbc", 33
}

/**
* 变量及数据类型
 */
func TestDataType() {
	// 批量声明变量
	var (
		name  string
		pAge  int
		isMan bool
	)

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", pAge)
	fmt.Printf("sex: %v\n", isMan)

	// 类型推导
	var pName = "zhangbc"
	fmt.Printf("name: %v\n", pName)

	// 短变量声明
	shortVarName := "shortVarName"
	fmt.Printf("shortVarName: %v\n", shortVarName)

	// 匿名变量
	name, _ = GetNameAndAge()
	fmt.Printf("name: %v\n", name)

	// 常量声明
	const PI float32 = 3.14159
	fmt.Printf("PI: %v\n", PI)

	// iota
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	const (
		a4 = iota
		_  // or b = 100
		a5 = iota
	)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)

	a := 100
	p := &a
	fmt.Printf("%T\n", p)

	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)
	fmt.Printf("arr: %v\n", arr)

	splitArr := []int{1, 2, 3}
	fmt.Printf("%T\n", splitArr)
	fmt.Printf("%T\n", f)

	// 布尔类型
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false
	b5 := true
	b6 := false
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

	age := 28
	condition := age >= 18
	if condition {
		fmt.Println("你已成年了.")
	} else {
		fmt.Println("你还未成年.")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}
}
