package fonction_go

import (
	"net/http"
)
type login struct{
	password_log	string
	mail_log	string
}

type register struct{
	password	string
	username		string
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
		password_log: r.FormValue("Password_log"),
		mail_log:	r.FormValue("Mail_log"),
	}
	return log;
}

func extractReg(r *http.Request) register{
	reg := register{
		password: r.FormValue("Password"),
		username:	r.FormValue("User"),
		mail:	r.FormValue("Mail"),
	}
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