package fonction_go

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	tmplPath := "login.html";
	path := "static/auth/";
	page:= extractPage(r);
	
	
	db, err := dbConnection()
	createUserTable(db)
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err);
		return;
	}
	if r.Method == http.MethodPost {
		switch page{
		case "log":
			println("log")
			params_log :=extractLog(r);
			l,err := logUser(db,params_log);
			if err != nil {
				handleError(w, "Erreur lors de la connexion", http.StatusInternalServerError, err);
				return;
			}
			if(l){
				path="static/protected/"
				tmplPath = "home.html"
			}

		default:
			params_reg :=extractReg(r);
			_,err := insertUser(db, params_reg);
			if err != nil {
				handleError(w, "Erreur lors de l'insertion dans la database", http.StatusInternalServerError, err);
				return;
			}
		}
	}

	
	defer db.Close();
	
	tmpl, err := template.New(tmplPath).ParseFiles(path + tmplPath);
	
	if err := tmpl.Execute(w,tmpl); err != nil {
		handleError(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError, err);
		return;
	}

}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode);
	if err != nil {
		log.Printf("%s: %v", message, err);
	}
}