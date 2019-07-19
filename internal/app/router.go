package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter assign route path
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("web/static")))
	r.PathPrefix("/static/").Handler(s)
	// Route path
	r.HandleFunc("/", indexView)
	r.HandleFunc("/admin/", adminView)
	r.HandleFunc("/admin/create/", createView)
	r.HandleFunc("/admin/{id}/", updateView)
	r.HandleFunc("/admin/delete/{id}/", deleteView)
	r.HandleFunc("/exercise/{id}/", exerView)
	// API path
	r.HandleFunc("/api/", createExer).Methods("POST")
	r.HandleFunc("/api/", getExers).Methods("GET")
	r.HandleFunc("/api/{id}/", getExer).Methods("GET")
	r.HandleFunc("/api/{id}/", updateExer).Methods("PUT")
	r.HandleFunc("/api/{id}/", deleteExer).Methods("DELETE")
	return r
}
