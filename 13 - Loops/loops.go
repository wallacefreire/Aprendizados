package main

import (
	"fmt"
)

func main() {
	nomes := [3]string{"Alice", "Bob", "Charlie"}

	for _, valor := range nomes {
		fmt.Println(valor)
	}
}
