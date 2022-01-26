package main

import (
	"fmt"
)

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
			// panic("second panic")
		}
	}()
	fmt.Println("Some useful wrok")
	panic("something bad happend")
}

func main() {
	deferTest()
}
