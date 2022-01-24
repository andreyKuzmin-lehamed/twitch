package main

import "fmt"

func main() {
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
