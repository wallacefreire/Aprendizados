package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := user{
			"João",
			"joao.pedro@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)

	})

	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
