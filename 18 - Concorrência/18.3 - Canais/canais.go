package main

import (
	"fmt"
	"time"
)

// Posso ter centenas de Goroutines rodando ao mesmo tempo...
// E usar os canais para o momento em que você precise que ela retorne um valor

func main() {
	// canal --- É uma comunicação, pode enviar ou receber dados.
	canal := make(chan string)
	go escrever("Hello World", canal)

	fmt.Println("Depois da função escrever() começar a ser executado")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa, até a próxima, pessoal!!!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Essa linha vai enviar o valor
		time.Sleep(time.Second)
	}
	close(canal)
}
