package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outro_numero := numero; outro_numero > 0 {
		fmt.Println("Número é maior que 0")
	} else {
		fmt.Println("Número é menor ou igual a 0")
	}
}
