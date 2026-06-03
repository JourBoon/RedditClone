package main

import (
	fonction_go "Fonction_go/backend"
	"log"
	"net/http"
)

func main() {
	fonction_go.RoutePages()

	log.Println("Serveur lancé sur http://localhost:4040")
	log.Fatal(http.ListenAndServe(":4040", nil))
}
