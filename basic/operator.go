package basic

import "fmt"

/**
* 运算符
 */
func TestOperator() {
	a := 100
	b := 20

	// 算术运算符
	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))

	a++
	b--
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// 关系运算符
	fmt.Printf("(a > b): %v\n", (a > b))
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a < b): %v\n", (a < b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b))
	fmt.Printf("(a != b): %v\n", (a != b))

	// 逻辑运算符
	b1 := true
	b2 := false
	fmt.Printf("(b1 && b2): %v\n", (b1 && b2))
	fmt.Printf("(b1 || b2): %v\n", (b1 || b2))
	fmt.Printf("(!b2): %v\n", (!b2))

	// 位运算符
	bit1 := 4
	bit2 := 8
	fmt.Printf("bit1: %b\tbit2: %b\n", bit1, bit2)
	fmt.Printf("(bit1 & bit2): %v\n", (bit1 & bit2))
	fmt.Printf("(bit1 | bit2): %v\n", (bit1 | bit2))
	fmt.Printf("(bit1 ^ bit2): %v\n", (bit1 ^ bit2))
	fmt.Printf("(bit1 >> 2): %v\n", (bit1 >> 2))
	fmt.Printf("(bit2 << 2): %v\n", (bit2 << 2))

	// 赋值运算符
	x := 100
	x = 200
	fmt.Printf("x: %v\n", x)
	x += 100
	fmt.Printf("x+=100: %v\n", x)
}
