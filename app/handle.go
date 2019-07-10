package app

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Declare templates
var tmlp = getTemplates()

// Context save data transfer to template
type Context map[string]string

func getTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func apiPOST(apiURL string, data url.Values) {

}

func indexView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Homepage"}
	tmlp.ExecuteTemplate(w, "index.html", ctx)
}

func adminView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Adminsite"}
	tmlp.ExecuteTemplate(w, "admin.html", ctx)
}

func createView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Create new exercise"}
	if r.Method == "POST" {
		apiURL := "http://localhost:8080/api/"
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()
		log.Println(handler.Header)
		contentFile, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}
		data := url.Values{}
		data.Set("title", r.Form.Get("title"))
		data.Set("description", r.Form.Get("description"))
		data.Set("testcase", string(contentFile))
		client := &http.Client{}
		rq, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
		rq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		resp, _ := client.Do(rq)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		resp.Body.Close()
		tmlp.ExecuteTemplate(w, "create.html", ctx)
	} else {
		tmlp.ExecuteTemplate(w, "create.html", ctx)
	}
}

func exerView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Exercise"}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
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
