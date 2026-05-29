package fonction_go

import (
	"net/http"
)

type queryParams struct {
	Password	string
	User		string
	searchQuery string
	searchType  string
	home       	string
	sortOrder   string
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		searchQuery: r.URL.Query().Get("Search"),
		searchType:  r.URL.Query().Get("SearchType"),
		home:       r.URL.Query().Get("Home"),
		sortOrder:   r.URL.Query().Get("sort"),
	}

	return params
}