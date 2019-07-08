package main

import (
	"html/template"
	"net/http"
)

func getTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	ctx := context {
		Title: "Homepage",
	}
	tmlp.ExecuteTemplate(w, "index.html", ctx)
}

func adminView(w http.ResponseWriter, r *http.Request) {
	ctx := context {
		Title: "Adminsite",
	}
	tmlp.ExecuteTemplate(w, "admin.html", ctx)
}
