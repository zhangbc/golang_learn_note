package basic

import (
	"bytes"
	"fmt"
	"strings"
)

// WebSite 结构体声明
type WebSite struct {
	Name string
}

/**
* 字符串类型
 */
func TestString() {
	var str1 string = "Hello world!"
	var str2 = "Hello World!"
	str3 := "Hello World!"
	str4 := `
	line1
	line2
	line3`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("str2: %v\n", str2)
	fmt.Printf("str3: %v\n", str3)
	fmt.Printf("str4: %v\n", str4)

	// 字符串连接
	var str12 string = str1 + str2
	fmt.Printf("str12: %v\n", str12)

	msg := fmt.Sprintf("%s, %s", str2, str3)
	fmt.Printf("msg: %v\n", msg)

	str_join := strings.Join([]string{str1, str4}, ",")
	fmt.Printf("str_join: %v\n", str_join)

	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	// 转义字符
	str5 := "Hello \n world"
	fmt.Printf("str5: %v\n", str5)

	// 字符串的切片操作
	str_split := "Hello World"
	start := 2
	end := 5
	fmt.Printf("str_split[start]: %c\n", str_split[start])
	fmt.Printf("str_split[start:end]: %v\n", str_split[start:end])
	fmt.Printf("str_split[start:]: %v\n", str_split[start:])

	//字符串函数
	fmt.Printf("len(str_split): %v\n", len(str_split))
	fmt.Printf("strings.Split(str_split, \" \"): %v\n", strings.Split(str_split, " "))
	fmt.Printf("strings.Contains(str_join, \"Hello\"): %v\n", strings.Contains(str_join, "Hello"))
	fmt.Printf("strings.ToLower(str_split): %v\n", strings.ToLower(str_split))
	fmt.Printf("strings.ToUpper(str_split): %v\n", strings.ToUpper(str_split))
	fmt.Printf("strings.HasPrefix(str_split, \"Hello\"): %v\n", strings.HasPrefix(str_split, "Hello"))
	fmt.Printf("strings.HasSuffix(str_split, \"world\"): %v\n", strings.HasSuffix(str_split, "world"))
	fmt.Printf("strings.Index(str_split, \"ll\"): %v\n", strings.Index(str_split, "ll"))

	// 格式化输出
	site := WebSite{Name: "douke360"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)

	bo := true
	fmt.Printf("bo: %t\n", bo)

	i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)
	i = 96
	fmt.Printf("i: %c\n", i)
	fmt.Printf("i: %x\n", 100)
	fmt.Printf("i: %x\n", 1234)

	x := 100
	p := &x
	fmt.Printf("p: %p\n", p)
}
