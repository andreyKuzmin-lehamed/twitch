package main

import (
	"fmt"
	"strconv"
)

func main() {
	const x, y = 3, 9
	fmt.Println(intCalc(x, y))
	fmt.Println(floatCalc(x, y))
}

func intCalc(x, y int) int {
	return x / y
}

func floatCalc(x, y float64) float64 {
	return x / y
}

func round() {
	fmt.Println("float32(0.4999999) result: ", float32(0.4999999))
	fmt.Println("float32(0.49999999) result: ", float32(0.49999999))
}

func untypedConstants() {

	fmt.Println("Объявляем нетипизированную константу со значением 2")
	const untypedConst = 2
	fmt.Println("Она становится нетипизированной, но int константой")
	fmt.Printf("const untypedConst = 2 // Тип: %T \n", untypedConst)

	fmt.Println("Можем обратиться с её помощью, к элементу слайса")
	var tmpSlice []string = []string{"a", "b", "c"}
	fmt.Println("В этом случае она воспринимается компилятором, как типизированная int константа")
	fmt.Println("tmpSlice: ", tmpSlice)
	fmt.Println("tmpSlice[untypedConst]: ", tmpSlice[untypedConst])

	fmt.Println("Можем объявить ещё одну нетипизированную int константу")
	const divider = 5
	fmt.Printf("const divider = 5 // Тип: %T \n", divider)
	fmt.Println("И поделим нашу старую константу на новую")
	fmt.Println("Если просто поделим untypedConst / divider, компилятор поделит как две типизированные int константы")
	fmt.Println("untypedConst / divider: ", untypedConst/divider)
	fmt.Println("Если явно приведем вторую константу к float64, то и первая станет типизированной float64 константой")
	fmt.Println("untypedConst / float64(divider): ", untypedConst/float64(divider))
}

func intToStr() {
	fmt.Println("Объявляем нетипизированную константу со значением 2")
	const untypedConst = 2
	fmt.Println("Она становится нетипизированной, но int константой")
	fmt.Println("Пытаемся, сконвертировать в строку")
	fmt.Println("string(untypedConst) result: ", string(untypedConst))
	fmt.Println("Но, получаем смайлик. Потому что в этом случае 2, воспринимается, как rune нужного символа.")
	fmt.Println("А 2 - это rune смайлика такого.")
	fmt.Println("А перевести число в строковый тип, можно вот так.")
	fmt.Println("strconv.Itoa(untypedConst) result: ", strconv.Itoa(untypedConst))
}
