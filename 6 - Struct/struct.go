package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var primeiro_usuario usuario
	primeiro_usuario.nome = "Wallace Freire"
	primeiro_usuario.idade = 21
	fmt.Println(primeiro_usuario)

	segundo_usuario := usuario{"Juliana Aveiro", 25}
	fmt.Println(segundo_usuario)

	terceiro_usuario := usuario{nome: "André Francisco"}
	fmt.Println(terceiro_usuario)

	quarto_usuario := usuario{"Daniel Daniele", 99}
	fmt.Println(quarto_usuario)
}
