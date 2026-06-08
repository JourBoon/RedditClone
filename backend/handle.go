package fonction_go

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func renderTemplateWithData(w http.ResponseWriter, tmpl string, data any) {
	path := filepath.Join("static", tmpl)

	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())
}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode)
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
