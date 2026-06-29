package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Conta Comigo!"))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nossos clientes da Conta Comigo!"))
}

func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação WEB

	// Cliente - Servidor

	// Request - Response

	// Rotas
	// URI - Identificador de Recurso
	// Método - GET, POST, PUST, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/user", user)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
