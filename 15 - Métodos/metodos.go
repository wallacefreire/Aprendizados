package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Débora Knupp", 24}
	usuario1.salvar()

	usuario2 := usuario{"Ademir Soares", 55}
	usuario2.salvar()

	maiordeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiordeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
