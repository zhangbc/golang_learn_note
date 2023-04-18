package basic

import "fmt"

// 数组的定义
func DefineArray() {
	var array1 [2]int
	var array2 [3]string
	var array3 = [2]bool{true, false}
	var array4 = [...]int{3, 4, 5}
	var array5 = [...]int{1: 2, 4: 100}
	fmt.Printf("array1: %v\n", array1)
	fmt.Printf("array2: %v\n", array2)
	fmt.Printf("array3: %v\n", array3)
	fmt.Printf("array4: %v\n", array4)
	fmt.Printf("array5: %v\n", array5)
}

// 数组的访问
func VisitArray() {
	var array = [...]int{100, 200, 3, 4, 5}
	fmt.Printf("array[0]: %v\n", array[0])
	array[0] = 2
	fmt.Printf("array[0]: %v\n", array[0])
	fmt.Printf("len(array): %v\n", len(array))

	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%v]: %v\t", i, array[i])
	}
	println()

	for i, v := range array {
		fmt.Printf("array[%v]: %v\t", i, v)
	}
	println()
}

// 切片的定义
func DefineSlice() {
	var slice1 []int
	var slice2 = make([]int, 5)
	var slice3 = []int{1, 2, 3, 4, 5}
	slice4 := slice3[3:]
	slice5 := slice3[2:5]
	slice6 := slice3[:]
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("slice4: %v\n", slice4)
	fmt.Printf("slice5: %v\n", slice5)
	fmt.Printf("slice6: %v\n", slice6)

	var sliceNull []int
	fmt.Printf("(sliceNull == nil): %v\n", (sliceNull == nil))
	fmt.Printf("len(sliceNull): %d, cap: %d\n", len(sliceNull), cap(sliceNull))
}

// 切片的访问
func VisitSlice() {
	var slice = []int{1, 2, 3, 45, 6}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]: %v\t", i, slice[i])
	}
	println()

	for i, v := range slice {
		fmt.Printf("slice[%v]: %v\t", i, v)
	}
	println()
}

// 切片的更新
func UpdateSlice() {
	// 添加元素
	var slice = []int{}
	slice = append(slice, 100)
	slice = append(slice, 200)
	slice = append(slice, 300)
	slice = append(slice, 400)
	fmt.Printf("slice: %v\n", slice)

	var temp = []int{1, 2, 3, 4}
	slice = append(slice, temp[:2]...)
	fmt.Printf("slice: %v\n", slice)

	// 删除元素
	slice = append(slice[:2], slice[3:]...)
	fmt.Printf("del slice[2]: %v\n", slice)

	// 更新元素
	slice[1] = 30000
	fmt.Printf("slice: %v\n", slice)
}

// 切片的查询
func QuerySlice() {
	var slice = []int{1, 2, 3, 4, 5}
	var key = 2
	for i, v := range slice {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

// 切片复制
func CopySlice() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1
	var slice3 = make([]int, 4)
	copy(slice3, slice1)
	slice2[0] = 1000
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
}

/**
* 数组 & 切片
 */
func TestArray() {
	// 数组定义
	println("=======数组定义=======")
	DefineArray()
	// 数组访问
	println("=======数组访问=======")
	VisitArray()
	// 切片定义
	println("=======切片定义=======")
	DefineSlice()
	// 切片访问
	println("=======切片访问=======")
	VisitSlice()
	// 切片更新
	println("=======切片更新=======")
	UpdateSlice()
	// 切片查询
	println("=======切片查询=======")
	QuerySlice()
	// 切片复制
	println("=======切片复制=======")
	CopySlice()
}
