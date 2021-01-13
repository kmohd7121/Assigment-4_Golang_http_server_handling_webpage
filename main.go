package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/index", indexhandler)
	r.HandleFunc("/end", homehandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "form.html", nil)
}
