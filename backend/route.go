package fonction_go

import (
	"fmt"
	"net/http"
)

func defaultPage() {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", RenderTemplate)
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", er)
		return
	}

	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	parms_log := extractLog(r);
	username := parms_log.username
	fmt.Printf("CSRF validate ;) Welcome to the forum %s", username)
}
