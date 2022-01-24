package main

import (
	"fmt"
)

func main() {
	fmt.Println("Обращаемся к символам строки, как к slice с байтами.")
	var str string = "Hello"
	fmt.Println("Вся строка:", str)
	fmt.Println("Байты через Printf:")
	fmt.Printf("% x\n", str)
	for i := 0; i < len(str); i++ {
		fmt.Println("Байт:", str[i], "Строковый тип:", string(str[i]))
	}

	fmt.Println("Но обращаясь к кириллице получаем проблемку.")
	var str_cir string = "Привет"
	fmt.Println("Вся строка:", str_cir)
	fmt.Println("Получая байты через Printf, видим, что их больше чем символов:")
	fmt.Printf("% x\n", str_cir)
	fmt.Println("Теперь немного изменяем Printf, чтобы получать юникод:")
	fmt.Printf("%+q\n", str_cir)
	for i := 0; i < len(str_cir); i++ {
		fmt.Println("Байт:", str_cir[i], "Строковый тип:", string(str_cir[i]))
	}

	fmt.Println("Конвертнем из []byte в []rune, и пробежися ещё раз:")
	str_cir_runes := []rune(str_cir)
	for i := 0; i < len(str_cir_runes); i++ {
		fmt.Println("Руна:", str_cir_runes[i], "Строковый тип:", string(str_cir_runes[i]))
	}

	fmt.Println("Переберём строку с кириллицей через rune, тут без проблем.")
	for _, r := range str_cir {
		fmt.Println("Руна:", r, "Строковый тип:", string(r))
	}
}

func strLiterals() {
	str := "🔴"
	// fmt.Println(str)
	byte_count := 0
	rune_count := 0
	length := len(str)

	for i := 0; i < len(str); i++ {
		fmt.Println("byte в byte виде:")
		fmt.Println(str[i])
	}

	for i := 0; i < len(str); i++ {
		fmt.Println("byte в строковом виде:")
		fmt.Println(string(str[i]))
		byte_count++
	}

	for _, el := range str {
		fmt.Println("rune в rune виде:")
		fmt.Println(el)
	}

	for _, el := range str {
		fmt.Println("rune в строковом виде:")
		fmt.Println(string(el))
		rune_count++
	}

	fmt.Println("Длина строки \"", str, "\" равняется ", length)
	fmt.Println("Отдельных byte: ", byte_count)
	fmt.Println("Отдельных rune: ", rune_count)
}
