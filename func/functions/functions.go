package main

import "fmt"

// Обычное объявление
func singleIn(in int) int {
	return in
}

// много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// именнованный результат
func nameReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error happend")
		// аналогично return rez, err
		// или return 1, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

// не фиксированное колличество параметров

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	fmt.Println(singleIn(1))
	fmt.Println(multIn(1, 2, 3))
	fmt.Println(nameReturn())
	fmt.Println(multipleReturn(3))
	fmt.Println(multipleNamedReturn(false))

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))

}
