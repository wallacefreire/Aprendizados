package main

import (
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Pingolin","raca":"Buldogue","idade":3}`

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		fmt.Println("Deu erro aqui po", err)
		return
	}
	fmt.Println(c)
}
