package fonction_go

import (
	"net/http"
)

type queryParams struct {
	password	string
	user		string
	mail		string
	searchQuery string
	searchType  string
	home       	string
	sortOrder   string
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		password: r.FormValue("Password"),
		user:	r.FormValue("User"),
		mail:r.FormValue("Mail"),
		searchQuery: r.FormValue("Search"),
		searchType:  r.FormValue("SearchType"),
		home:       r.FormValue("Home"),
		sortOrder:   r.FormValue("sort"),
	}

	return params
}