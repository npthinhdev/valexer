package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createExer(w http.ResponseWriter, r *http.Request) {
	exer := Exercise{}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exer.Title = r.Form.Get("title")
	exer.Description = r.Form.Get("description")
	exer.Testcase = r.Form.Get("testcase")
	err = createDBExer(&exer)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, "Create success exercise")
}

func getExers(w http.ResponseWriter, r *http.Request) {
	exers, err := getDBExers()
	if err != nil {
		log.Println(err)
	}
	exersByte, err := json.Marshal(exers)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(exersByte)
}

func getExer(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer, err := getDBExer(id)
	if err != nil {
		log.Println(err)
	}
	exerByte, err := json.Marshal(exer)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(exerByte)
}

func updateExer(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := Exercise{}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exer.Title = r.Form.Get("title")
	exer.Description = r.Form.Get("description")
	exer.Testcase = r.Form.Get("testcase")
	err = updateDBExer(id, &exer)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, "Update success exercise")
}

func deleteExer(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("Url id is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := deleteDBExer(id)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, "Delete success exercise")
}
