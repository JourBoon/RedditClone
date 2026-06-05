package fonction_go

import (
	"net/http"
)

type login struct {
	password     string
	mail         string
	username	string
	sessionToken string
	CsrfToken    string
}

type register struct {
	password string
	username string
	mail     string
}

type message struct{
	subject		string
	tags			[]string
	body 		string
}

type queryParams struct {
	searchQuery string
	searchType  string
	home        string
	sortOrder   string
}

func extractPage(r *http.Request) string{
	return r.FormValue("page");
}

func extractLog(r *http.Request) login {
	var username string;
	query := `SELECT username FROM users WHERE email=(?)`
	
	db,_ := dbConnection()
	
	err := db.QueryRow(query, r.FormValue("mail")).Scan(&username)
	
	if err != nil {
		print("Erreur lors de la récupération du username")
	}
	println("extract",username,r.FormValue("mail"))
	log := login{
		password:     r.FormValue("password"),
		mail:         r.FormValue("mail"),
		username: 	  username,
		sessionToken: "",
		CsrfToken:    "",
	}
	return log;
}

func extractMess(r *http.Request) message{

	mess := message{
		subject: r.FormValue("subject"),
		tags:	r.PostForm["tag"],
		body:	r.FormValue("body"),
	}
	return mess;
}

func extractReg(r *http.Request) register{
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
