package fonction_go

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type login struct {
	password     string
	mail         string
	username     string
	sessionToken string
	CsrfToken    string
}

type register struct {
	password string
	username string
	mail     string
}

type message struct {
	subject string
	tags    []string
	body    string
}

type queryParams struct {
	searchQuery string
	searchType  string
	home        string
	sortOrder   string
}

func extractPage(r *http.Request) string {
	return r.FormValue("page")
}

func extractLog(r *http.Request) login {
	var username string
	query := `SELECT username FROM users WHERE email=(?)`

	db, err := dbConnection()
	if err != nil {
		fmt.Println("Erreur DB dans extractLog:", err)
		return login{
			password:     r.FormValue("password"),
			mail:         r.FormValue("mail"),
			username:     "",
			sessionToken: "",
			CsrfToken:    "",
		}
	}
	defer db.Close()

	err = db.QueryRow(query, r.FormValue("mail")).Scan(&username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Erreur lors de la récupération du username:", err)
		}
	}
	println("extract", username, r.FormValue("mail"))
	log := login{
		password:     r.FormValue("password"),
		mail:         r.FormValue("mail"),
		username:     username,
		sessionToken: "",
		CsrfToken:    "",
	}
	return log
}

func extractMess(r *http.Request) message {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Erreur ParseForm:", err)
	}

	tagValues := r.Form["tag"]
	if len(tagValues) == 0 {
		tagValues = r.Form["tag[]"]
	}

	mess := message{
		subject: r.FormValue("subject"),
		tags:    tagValues,
		body:    r.FormValue("body"),
	}
	return mess
}

func extractReg(r *http.Request) register {
	println("extraction...")
	reg := register{
		password: r.FormValue("password"),
		username: r.FormValue("user"),
		mail:     r.FormValue("mail"),
	}
	return reg
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		searchQuery: r.FormValue("search"),
		searchType:  r.FormValue("searchType"),
		home:        r.FormValue("home"),
		sortOrder:   r.FormValue("sort"),
	}
	return params
}
