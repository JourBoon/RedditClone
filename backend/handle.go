package fonction_go

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	tmplPath := "login.html"
	path := "static/auth/"
	page := extractPage(r)

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	if r.Method == http.MethodPost {
		switch page {
		case "log":
			params_log := extractLog(r)
			l, err := logUser(db, params_log)
			if err != nil {
				handleError(w, "Erreur lors de la connexion", http.StatusInternalServerError, err)
				return
			}
			if l {
				// Gestion de la session
				params_log := extractLog(r)
				fmt.Printf("test")

				params_log.sessionToken = generateToken(32)
				params_log.csrfToken = generateToken(32)

				http.SetCookie(w, &http.Cookie{
					Name:     "session_token",
					Value:    params_log.sessionToken,
					Expires:  time.Now().Add(24 * time.Hour),
					HttpOnly: true,
				})

				http.SetCookie(w, &http.Cookie{
					Name:     "csrf_token",
					Value:    params_log.csrfToken,
					Expires:  time.Now().Add(24 * time.Hour),
					HttpOnly: false,
				})

				addSessionToken(db, params_log)
				addCsrfToken(db, params_log)

				if err != nil {
					handleError(w, "Erreur lors de l'insertion des token dans la db", http.StatusInternalServerError, err)
					return
				}

				path = "static/protected/"
				tmplPath = "home.html"
			}

		case "message":
			params_mess := extractMess(r);
			postMess(db,params_mess.subject,params_mess.body);
		default:
			params_reg := extractReg(r)
			_, err := insertUser(db, params_reg)
			if err != nil {
				handleError(w, "Erreur lors de l'insertion dans la database", http.StatusInternalServerError, err)
				return
			}
		}
	}

	defer db.Close()

	tmpl, err := template.New(tmplPath).ParseFiles(path + tmplPath)

	if err := tmpl.Execute(w, tmpl); err != nil {
		handleError(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError, err)
		return
	}
}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode)
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
