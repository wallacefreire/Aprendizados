package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 01")
}

func funcao2() {
	fmt.Println("Executando a função 02")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função, verificando se o aluno foi aprovado...")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}
func main() {
	fmt.Println(alunoAprovado(4, 10))
}
