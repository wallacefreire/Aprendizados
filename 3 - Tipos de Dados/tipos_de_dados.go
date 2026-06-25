package main

import "fmt"

func main() {
	numero := -100
	fmt.Println(numero)

	var numero2 uint32 = 5
	fmt.Println(numero2)

	// alias
	var numero3 rune = 12345
	fmt.Println(numero3)

	// Byte = uint8
	var numero4 byte = 10
	fmt.Println(numero4)

	var numeroReal float32 = 543.90
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 42132.09
	fmt.Println(numeroReal2)

	// Texto

	var str string = "Texto Exemplo"
	fmt.Println(str)

	str2 := "Texto exemplo 2"
	fmt.Println(str2)

	// Fim Texto
}
