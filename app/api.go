package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Fprintln(w, "Success")
}

func getExers(w http.ResponseWriter, r *http.Request) {
	exers, err := getDBExers()
	if err != nil {
		log.Println(err)
	}
	exerBytes, err := json.Marshal(exers)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(exerBytes)
}
