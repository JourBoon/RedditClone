package fonction_go

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func renderTemplateWithData(w http.ResponseWriter, tmpl string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := filepath.Join("static", tmpl)

	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode)
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
