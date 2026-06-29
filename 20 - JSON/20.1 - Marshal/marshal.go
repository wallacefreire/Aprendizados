package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Pitoco", "Vira-Lata", 14}
	fmt.Println(c)

	cachorroEmJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Não consegui abrir o arquivo:", err)
		return
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Pitico",
		"raca": "vira vira",
	}

	cachorro2EmJSON, err := json.Marshal(c2)
	if err != nil {
		fmt.Println("Não consegui abrir o arquivo:", err)
		return
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
