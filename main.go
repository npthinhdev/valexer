package main

import (
	"net/http"
)

// Declare templates
var tmlp = getTemplates()

type context struct {
	Title string
}

func main() {
	r := getRouter()
	http.ListenAndServe(":8080", r)
}
