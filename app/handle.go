package app

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Declare templates
var tmlp = getTemplates()

type context map[string]string

func getTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func indexView(w http.ResponseWriter, r *http.Request) {
	ctx := context{"Title": "Homepage"}
	tmlp.ExecuteTemplate(w, "index.html", ctx)
}

func adminView(w http.ResponseWriter, r *http.Request) {
	ctx := context{"Title": "Adminsite"}
	tmlp.ExecuteTemplate(w, "admin.html", ctx)
}

func exerView(w http.ResponseWriter, r *http.Request) {
	ctx := context{"Title": "Exercise"}
	if r.Method == "POST" {
		r.ParseMultipartForm(1 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		log.Println(handler.Header)
		contentFile, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}
		ctx["Code"] = string(contentFile)
		tmlp.ExecuteTemplate(w, "exercise.html", ctx)
	} else {
		tmlp.ExecuteTemplate(w, "exercise.html", ctx)
	}
}
