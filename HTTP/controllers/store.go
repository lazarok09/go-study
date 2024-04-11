package controllers

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func Store(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates.ExecuteTemplate(w, "index.html", nil)
}
