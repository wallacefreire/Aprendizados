package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/funcionarios", servidor.CriarFuncionario).Methods(http.MethodPost)
	router.HandleFunc("/funcionarios", servidor.BuscarFuncionarios).Methods(http.MethodGet)
	router.HandleFunc("/funcionarios/{id}", servidor.BuscarFuncionario).Methods(http.MethodGet)
	router.HandleFunc("/funcionarios/{id}", servidor.AtualizarFuncionario).Methods(http.MethodPut)

	fmt.Println("Seja bem-vindo a porta 5k")
	log.Fatal(http.ListenAndServe(":5000", router))
}
