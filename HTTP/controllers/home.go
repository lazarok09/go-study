package controllers

import (
	"net/http"
	"text/template"
)

var templates *template.Template

type User struct {
	Name string
	Age  int
}
type MetaData struct {
	Title       string
	Description string
	User        User
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("views/*.html"))
	meta := MetaData{Title: "Welcome to my page",
		Description: "People used to say that feelings are not the best indicative of our wareness",
		User:        User{Name: "Lázaro Souza", Age: 24}}

	templates.ExecuteTemplate(w, "index.html", meta)
}
