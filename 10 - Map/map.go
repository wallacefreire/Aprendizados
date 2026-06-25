package main

import "fmt"

func main() {
	fmt.Println("Map")

	usuario := map[string]string{
		"nome":      "Micael",
		"sobrenome": "Freire",
	}

	fmt.Println(usuario["nome"])
}
