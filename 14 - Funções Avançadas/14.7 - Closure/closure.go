package main

import "fmt"

func closure() func() {
	texto := "Dentro da Closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "Dentro da Main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
