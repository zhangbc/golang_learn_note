package basic

import (
	"fmt"
)

// 判断奇偶性
func JudeOddEven() {
	var num int
	fmt.Println("Please a nmber:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("This is an even num: %v\n", num)
	} else {
		fmt.Printf("This is an odd num: %v\n", num)
	}
}

// 分数等级
func GetScoreDegree(score float32) {
	if score >= 60 && score < 70 {
		fmt.Printf("\"C\": %v\n", "C")
	} else if score >= 70 && score < 90 {
		fmt.Printf("\"B\": %v\n", "B")
	} else if score >= 90 {
		fmt.Printf("\"A\": %v\n", "A")
	} else {
		fmt.Printf("\"D\": %v\n", "D")
	}
}

// 分数等级-switch
func GetScoreSwitch() {
	grade := 'D'
	switch grade {
	case 'A':
		fmt.Printf("'A': %v\n", "优秀")
	case 'B':
		fmt.Printf("'B': %v\n", "良好")
	case 'C':
		fmt.Printf("'C': %v\n", "合格")
	default:
		fmt.Printf("'D': %v\n", "不合格")
	}
}

// 判断工作日
func IsWorkDay() {
	day := 2
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Printf("workday: %v\n", day)
		fallthrough
	case 6, 7:
		fmt.Printf("holiday day: %v\n", day)
	default:
		fmt.Printf("Not a day: %v\n", day)
	}
}

// 判断三个数的大小
func CompareThreeNumbers() {
	a, b, c := 1, 2, 3
	if a > b {
		if a > c {
			fmt.Printf("The max number is: %v\n", a)
		}
	} else {
		if b > c {
			fmt.Printf("The max number is: %v\n", b)
		} else {
			fmt.Printf("The max number is: %v\n", c)
		}
	}
}

// 使用标签 -- break
func BreakLabel() {
MYLABEL:
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
		if i > 5 {
			break MYLABEL
		}
	}
	println("break MYLABEL...")
}

// 使用标签 -- for
func ForLabel() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nrow: %v\n", i)
	FORLABL:
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				continue FORLABL
			}
			fmt.Printf("(%v,%v) ", i, j)
		}
	}
	println()
}

// 使用标签 -- goto
func GotoLabel() {
	v := 0
	if v == 0 {
		goto GOTOLABEL
	} else {
		println("other...")
	}

GOTOLABEL:
	println("next...")
}

// 使用标签 -- goto
func GotoForLabel() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nrow: %v\n", i)
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto ENDLABEL
			}
			fmt.Printf("(%v,%v) ", i, j)
		}
	}
	println()

ENDLABEL:
	println("end...")
}

/**
* 流程控制
 */
func TestController() {
	// if语句
	a := 1
	b := 2
	if a > b {
		fmt.Printf("a: %v\n", a)
	} else {
		fmt.Printf("b: %v\n", b)
	}

	if age := 20; age > 18 {
		fmt.Printf("You are adult, age: %v\n", age)
	} else {
		fmt.Printf("You are not adult, age: %v\n", age)
	}

	// if else 语句
	var name string
	var age int
	var email string
	fmt.Println("Please input[name, age, emaill]")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v,\tage: %v,\temail: %v\n",
		name, age, email)

	JudeOddEven()

	GetScoreDegree(78)

	CompareThreeNumbers()

	// switch语句
	GetScoreSwitch()

	IsWorkDay()

	// for 循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	j := 0
	for j < 10 {
		fmt.Printf("j: %v\n", j)
		j++
	}

	// 永真循环
	// for {
	// 	fmt.Printf("for: %v\n", "run....")
	// }

	// for range 循环
	var array = [...]int{1, 2, 3, 4, 5}
	fmt.Print("\nArray: ")
	for i, v := range array {
		fmt.Printf("i: %d, v: %d ", i, v)
	}

	var sp = []int{1, 2, 3}
	fmt.Print("\nSplit array: ")
	for _, v := range sp {
		fmt.Printf("%v ", v)
	}

	var mapItems = make(map[string]string, 0)
	mapItems["name"] = "zhangbc"
	mapItems["age"] = "20"
	mapItems["email"] = "zhangbocheng189@163.com"
	fmt.Print("\nMap k(v): ")
	for k, v := range mapItems {
		fmt.Printf("%v: %v ", k, v)
	}

	s := "hello world"
	fmt.Print("\nString: ")
	for _, v := range s {
		fmt.Printf("%c", v)
	}

	// break 语句
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("break...")
			break
		}
	}

	i := 2
	switch i {
	case 1:
		println("case 1...")
	case 2:
		println("case 2...")
		fallthrough
	case 3:
		println("case 3...")
	default:
		println("default...")
	}

	BreakLabel()

	// continue 语句
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v ", i)
		} else {
			continue
		}
	}
	println()

	ForLabel()

	// goto 语句
	GotoLabel()
	GotoForLabel()
}
