package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var testeWaitGroup sync.WaitGroup

	testeWaitGroup.Add(2) // 2 Goroutines para serem executadas

	go func() { // Por essa função anônima estar sendo chamada com a cláusula go na frente...
		escrever("Hello World")
		testeWaitGroup.Done() // O trabalho está feito, retirar um Goroutine (-1)
	}()

	go func() { // Ela vai ser rodada de forma concorrente.
		escrever("Wallace Freire || $22.279.211,89")
		testeWaitGroup.Done() // O trabalho está feito, retirar um Goroutine (-1)
	}()

	testeWaitGroup.Wait() // Esperar a contagem das Goroutines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
