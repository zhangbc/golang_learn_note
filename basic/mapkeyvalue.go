package basic

import "fmt"

// Map定义
func DefineMap() {
	// 类型的声明
	var map1 map[string]string
	map1 = make(map[string]string)
	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map1: %T\n", map1)

	// 初始化
	var map2 = map[string]string{"name": "zhangbc", "age": "20"}
	fmt.Printf("map2: %v\n", map2)

	map3 := make(map[string]string)
	map3["name"] = "wangw"
	fmt.Printf("map3: %v\n", map3)
}

// Map基本操作
func OperatorMap() {
	var map1 = map[string]string{"name": "zhangbc", "age": "20"}
	var key1 = "name"
	var key2 = "age1"
	fmt.Printf("map1[key1]: %v\n", map1[key1])
	v, exist := map1[key1]
	fmt.Printf("(k, v): %v-%v\n", exist, v)

	println("====================")

	v, exist = map1[key2]
	fmt.Printf("(k, v): %v-%v\n", exist, v)
}

// Map遍历操作
func VisitMap() {
	var map1 = map[string]string{"name": "zhangbc", "age": "20"}
	for k := range map1 {
		fmt.Printf("k: %v\n", k)
	}
	println("====================")
	for k, v := range map1 {
		fmt.Printf("(k, v): %v-%v\n", k, v)
	}
	println("====================")
}

/**
* Map
 */
func TestMap() {
	// Map定义
	println("=======Map定义=======")
	DefineMap()
	// Map基本操作
	println("=======Map基本操作=======")
	OperatorMap()
	// Map遍历操作
	println("=======Map遍历操作=======")
	VisitMap()
}
