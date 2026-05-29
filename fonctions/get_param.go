package fonction_go

import (
	"net/http"
)

type queryParams struct {
	password	string
	password_log	string
	user		string
	mail		string
	mail_log	string
	searchQuery string
	searchType  string
	home       	string
	sortOrder   string
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		password: r.FormValue("Password"),
		password_log: r.FormValue("Password_log"),
		user:	r.FormValue("User"),
		mail:	r.FormValue("Mail"),
		mail_log:	r.FormValue("Mail_log"),
		searchQuery: r.FormValue("Search"),
		searchType:  r.FormValue("SearchType"),
		home:       r.FormValue("Home"),
		sortOrder:   r.FormValue("sort"),
	}

	return params
}