package main

import (
	fonction_go "Fonction_go/backend"
	"log"
	"net/http"
)

func main() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", fonction_go.RenderTemplate)

	log.Println("Serveur lancé sur http://localhost:4040")
	log.Fatal(http.ListenAndServe(":4040", nil))
}