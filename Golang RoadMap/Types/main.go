package main

import (
	"fmt"
)

func main() {
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
