package fonction_go

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {

	tmplPath := "login.html";
	dbConnection()
	
	params:= extractQueryParams(r);
	tmpl, err := template.New(tmplPath).ParseFiles("static/auth/" + tmplPath);
	if err != nil {
		handleError(w, "Erreur lors du chargement du template", http.StatusInternalServerError, err);
		return;
	}

	if err := tmpl.Execute(w,tmpl); err != nil {
		handleError(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError, err);
		return;
	}

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
    insertUser(db, params)
	defer db.Close()

	fmt.Println(params.mail);
	//fmt.Println(params.password)
	//fmt.Println(hashedPassword(params.password))
}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode);
	if err != nil {
		log.Printf("%s: %v", message, err);
	}
}