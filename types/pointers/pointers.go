package main

import "fmt"

func main() {
	a := 2
	b := &a

	*b = 3 // a = 3
	fmt.Println(*b)
	c := &a // новый указатель на переменную a

	// получение указателя на переменную типа int
	// инициализированно значением по-умолчанию
	d := new(int)
	*d = 12
	fmt.Println(*d)
	*c = *d // c = 12 -> a = 12
	fmt.Println(*c)
	*d = 13 // c && a не изменились
	fmt.Println(*d)
	fmt.Println(*c)
	fmt.Println(a)

	c = d   // тереоь c указывает туда же, куда и d
	*c = 14 // c = 14 -> d = 14, a = 12
	fmt.Println(*c)
	fmt.Println(*d)

}
