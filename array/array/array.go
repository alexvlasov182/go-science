package main

import "fmt"

func main() {
	// размер массива является частью его типа

	// инициализация ззначениями по умолчанию
	var a1 [3]int // [0,0,0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)
}
