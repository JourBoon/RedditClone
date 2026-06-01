package fonction_go

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/text/cases"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	tmplPath := "login.html";
	
	page:= extractPage(r);
	
	tmpl, err := template.New(tmplPath).ParseFiles("static/auth/" + tmplPath);
	
	if err := tmpl.Execute(w,tmpl); err != nil {
		handleError(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError, err);
		return;
	}

	db, err := dbConnection()

	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err);
		return;
	}

	switch page{
	case "log":
		params_log :=extractLog(r);
		logUser(db,params_log);
	default:
		params_reg :=extractReg(r);
		insertUser(db, params_reg);
	}
	if err != nil {
		handleError(w, "Erreur lors du chargement du template", http.StatusInternalServerError, err);
		return;
	}

	
	defer db.Close();

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