package fonction_go

import (
	"fmt"
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
type queryParams struct {
	searchQuery string
	searchType  string
	home       	string
	sortOrder   string
}

func extractPage(r *http.Request) string{
	return r.FormValue("Page");
}

func extractLog(r *http.Request) login{
	log := login{
		password: r.FormValue("Password_log"),
		mail:	r.FormValue("Mail_log"),
	}
	return log;
}

func extractReg(r *http.Request) register{
	println("extraction...")
	reg := register{
		password: r.FormValue("Password"),
		username:	r.FormValue("User"),
		mail:	r.FormValue("Mail"),
	}
	fmt.Println(reg)
	return reg;
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		searchQuery: r.FormValue("Search"),
		searchType:  r.FormValue("SearchType"),
		home:       r.FormValue("Home"),
		sortOrder:   r.FormValue("sort"),
	}
	return params;
}