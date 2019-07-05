package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func getTemplates() *template.Template {
	var allFiles []string
	files, err := ioutil.ReadDir("templates")
	if err != nil {
		return nil
	}
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allFiles = append(allFiles, "templates/"+fileName)
		}
	}
	return template.Must(template.ParseFiles(allFiles...))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	tmlp.ExecuteTemplate(w, "index.html", nil)
}

func adminView(w http.ResponseWriter, r *http.Request) {
	tmlp.ExecuteTemplate(w, "admin.html", nil)
}
