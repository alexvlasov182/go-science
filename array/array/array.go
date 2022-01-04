package main

import "fmt"

func main() {
	// размер массива является частью его типа

	// инициализация ззначениями по умолчанию
	var a1 [3]int // [0,0,0]

	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Printf("a2 full %#v\n", a2)
	// определение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	// проверка при компиляции или при выполнении
	// invalid array index 4 (out of bounds for 3-element array)
	// a3[4] = 12
}
