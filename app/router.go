package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetRouter assign route path
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	// Static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(s)
	// Route path
	r.HandleFunc("/", indexView)
	r.HandleFunc("/admin/", adminView)
	r.HandleFunc("/admin/create/", createView)
	r.HandleFunc("/admin/{id}/", updateView)
	r.HandleFunc("/exercise/{id}/", exerView)
	// API path
	r.HandleFunc("/api/", createExer).Methods("POST")
	r.HandleFunc("/api/", getExers).Methods("GET")
	r.HandleFunc("/api/{id}/", getExer).Methods("GET")
	r.HandleFunc("/api/{id}/", updateExer).Methods("PUT")
	return r
}
