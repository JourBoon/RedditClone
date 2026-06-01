package fonction_go

import (
	"net/http"
)
type login struct{
	password	string
	mail		string
}

type register struct{
	password	string
	username	string
	mail		string
}

type message struct{
	subject		string
	body 		string
}

type queryParams struct {
	searchQuery string
	searchType  string
	home       	string
	sortOrder   string
}


func extractPage(r *http.Request) string{
	return r.FormValue("page");
}

func extractLog(r *http.Request) login{
	log := login{
		password: r.FormValue("password"),
		mail:	r.FormValue("mail"),
	}
	return log;
}

func extractMess(r *http.Request) message{
	mess := message{
		subject: r.FormValue("subject"),
		body:	r.FormValue("body"),
	}
	return mess;
}

func extractReg(r *http.Request) register{
	println("extraction...")
	reg := register{
		password: r.FormValue("password"),
		username:	r.FormValue("user"),
		mail:	r.FormValue("mail"),
	}
	return reg;
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		searchQuery: r.FormValue("search"),
		searchType:  r.FormValue("searchType"),
		home:       r.FormValue("home"),
		sortOrder:   r.FormValue("sort"),
	}
	return params;
}