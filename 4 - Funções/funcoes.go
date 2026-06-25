package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calculosMatematicos(integrante1, integrante2 int32) (int32, int32) {
	soma := integrante1 + integrante2
	subtracao := integrante1 - integrante2
	return soma, subtracao
}
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Função da Função")
	fmt.Println(resultado)

	determinacaoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(determinacaoSoma)
}
