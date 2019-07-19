package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

// Declare templates
var tmlp = getTemplates()

// Context save data transfer to template
type Context map[string]interface{}

func getTemplates() *template.Template {
	return template.Must(template.ParseGlob("web/template/*.html"))
}

func apiPOST(apiURL string, data url.Values) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func apiPUT(apiURL string, data url.Values) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("PUT", apiURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func apiDELETE(apiURL string) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("DELETE", apiURL, nil)
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func apiGET(apiURL string) []byte {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", apiURL, nil)
	r.Header.Add("Accept", "application/json")
	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func formFile(r *http.Request) ([]byte, error) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()
	log.Println(handler.Header)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return content, nil
}

func getTesting(id string, solution []byte) string {
	head := `package main
func main() {
	matchAll := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{F: TestSolution}}
	testing.Main(matchAll, tests, nil, nil)
}`
	apiURL := fmt.Sprintf("http://localhost:8080/api/%s/", id)
	body := apiGET(apiURL)
	exer := Exercise{}
	_ = json.Unmarshal(body, &exer)
	testcase := exer.Testcase
	mess := testcase + string(solution)
	re := regexp.MustCompile(`(package.*)|(import.*")|(import(?s).*"\n\))`)
	mess = string(head) + re.ReplaceAllString(mess, "")
	getFormat(&mess)
	result := getCompile(&mess)
	return result
}

func getFormat(mess *string) {
	apiURL := "https://play.golang.org/fmt"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("imports", "true")
	body := apiPOST(apiURL, data)
	var msg struct {
		Body string `json:"body"`
	}
	_ = json.Unmarshal(body, &msg)
	*mess = msg.Body
}

func getCompile(mess *string) string {
	apiURL := "https://play.golang.org/compile"
	data := url.Values{}
	data.Set("body", *mess)
	data.Set("version", "2")
	data.Set("withVet", "true")
	body := apiPOST(apiURL, data)
	var msg struct {
		Errors string
		Events []struct {
			Message string `json:"message"`
		}
	}
	_ = json.Unmarshal(body, &msg)
	if msg.Errors != "" {
		return msg.Errors
	}
	return msg.Events[0].Message
}

// Handle requets

func indexView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Homepage"}
	apiURL := "http://localhost:8080/api/"
	body := apiGET(apiURL)
	exers := []Exercise{}
	_ = json.Unmarshal(body, &exers)
	ctx["Exers"] = exers
	tmlp.ExecuteTemplate(w, "index.html", ctx)
}

func adminView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Adminsite"}
	apiURL := "http://localhost:8080/api/"
	body := apiGET(apiURL)
	exers := []Exercise{}
	_ = json.Unmarshal(body, &exers)
	ctx["Exers"] = exers
	tmlp.ExecuteTemplate(w, "admin.html", ctx)
}

func createView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Create new exercise"}
	if r.Method == "POST" {
		apiURL := "http://localhost:8080/api/"
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		fileContent, err := formFile(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := url.Values{}
		data.Set("title", r.Form.Get("title"))
		data.Set("description", r.Form.Get("description"))
		data.Set("testcase", string(fileContent))
		body := apiPOST(apiURL, data)
		log.Println(string(body))
		http.Redirect(w, r, "/admin/", http.StatusFound)
	} else {
		tmlp.ExecuteTemplate(w, "create.html", ctx)
	}
}

func exerView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Exercise"}
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/%s/", id)
	body := apiGET(apiURL)
	if string(body) == "null" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := Exercise{}
	_ = json.Unmarshal(body, &exer)
	ctx["Exer"] = exer
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		fileContent, err := formFile(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result := getTesting(id, fileContent)
		ctx["Result"] = string(result)
		tmlp.ExecuteTemplate(w, "exercise.html", ctx)
	} else {
		tmlp.ExecuteTemplate(w, "exercise.html", ctx)
	}
}

func updateView(w http.ResponseWriter, r *http.Request) {
	ctx := Context{"Title": "Update exercise"}
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/%s/", id)
	body := apiGET(apiURL)
	if string(body) == "null" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := Exercise{}
	_ = json.Unmarshal(body, &exer)
	ctx["Exer"] = exer
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		data := url.Values{}
		data.Set("title", r.Form.Get("title"))
		data.Set("description", r.Form.Get("description"))
		data.Set("testcase", r.Form.Get("testcase"))
		body := apiPUT(apiURL, data)
		log.Println(string(body))
		http.Redirect(w, r, "/admin/", http.StatusFound)
	} else {
		tmlp.ExecuteTemplate(w, "update.html", ctx)
	}
}

func deleteView(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/%s/", id)
	_ = apiDELETE(apiURL)
	http.Redirect(w, r, "/admin/", http.StatusFound)
}
