package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo

	// Paralelismo são 2 ou + tarefas sendo executadas ao mesmo tempo.
	// Para isso processador com mais de um núcleo.

	go escrever("Hello World")
	escrever("Wallace Freire || $22.279.211,89")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
