package basic

import (
	"fmt"
	"math"
	"unsafe"
)

/**
* 数字类型
 */
func TestNumber() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i=%v\n", i)
	}

	// 数字类型
	var i8 int
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var f32 float32
	var f64 float64
	var ui uint
	var imax, imin int
	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)
	ui = uint(math.MaxUint64)
	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)
	imax = int(math.MaxInt64)
	imin = int(math.MinInt64)
	fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)

	// 进制数
	var bit int = 10
	fmt.Printf("bit: %d\n", bit)
	fmt.Printf("bit: %b\n", bit)

	var ebit int = 077
	fmt.Printf("ebit: %o\n", ebit)

	var hbit int = 0xff
	fmt.Printf("hbit: %x\n", hbit)
	fmt.Printf("hbit: %X\n", hbit)

	// 浮点数
	fmt.Printf("math.Pi: %f\n", math.Pi)
	fmt.Printf("math.Pi: %.2f\n", math.Pi)

	// 复数
	var v1 complex64 = 1 + 2i
	var v2 complex128 = 2 + 3i
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
}
