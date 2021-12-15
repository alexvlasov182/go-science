package main

func main() {
	// пустая строка по-умолчанию
	var str string

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символами
	var world string = `Мир\n\t`

	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"
	hi := "some chinas symbols"

	// одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) for UTF-8 symbols
}
