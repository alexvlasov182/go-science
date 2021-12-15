package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// пустая строка по-умолчанию
	var str string
	fmt.Println(str)

	// со спец символами
	var hello string = "Привет\n\t"
	fmt.Println(hello)

	// без спец символами
	var world string = `Мир\n\t`
	fmt.Println(world)

	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"
	hi := "康熙字典体"
	fmt.Println(hi, helloWorld)

	// одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'
	fmt.Println(rawBinary)

	// rune (uint32) for UTF-8 symbols
	var someChines rune = '康'
	fmt.Println(someChines)

	// helloWorld := "Привет мир"
	fmt.Println(helloWorld)

	// конкатинация строк
	andGoodMorning := helloWorld + " и доброе утро"
	fmt.Println(andGoodMorning)

	// строки неизменяемы
	// cannot assign to helloWorld[0]
	// helloWorld[0] = 72
	fmt.Println(helloWorld)

	// получение длины строки
	byteLen := len(helloWorld)                    // 19 byte
	symbols := utf8.RuneCountInString(helloWorld) // 10 rune
	fmt.Println(byteLen, symbols)

	// получение подстроки, в байтах, не символах!
	// hello := helloWorld[:12] // Привет, 0-11 байты
	H := helloWorld[0] // byte 72 not "П"
	fmt.Println(H, hello)

	// конвертация в слайс байт и обратно
	// byteString = []byte(helloWorld)
	// helloWorld = string(byteString)

}
