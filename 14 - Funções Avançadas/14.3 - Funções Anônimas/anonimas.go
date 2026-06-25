package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 68)
	}("Passando os Parâmetros")

	fmt.Println(retorno)
}
