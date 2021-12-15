package main

import "fmt"

func main() {
	// int - платформо зависимый тип, 32/64
	var i int = 10

	// автоматически выбранный тип int
	var autoInt = -10

	//int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// платформо зависимый тип, 32/64
	var unsignedInt uint = 100500

	//uint8, uint16, uint32, uint64
	var unsingBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt, unsingBigInt)

	// float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	// bool
	var b bool // flase по-умолчанию
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	// complex64 complex128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i
	fmt.Println(c, c2)

}
