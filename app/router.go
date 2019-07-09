package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	// Static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(s)
	// Route path
	r.HandleFunc("/", indexView)
	r.HandleFunc("/admin/", adminView)
	r.HandleFunc("/exercise/", exerView)
	return r
}
