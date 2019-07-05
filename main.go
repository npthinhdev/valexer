package main

import (
	"net/http"
)

// Declare templates
var tmlp = getTemplates()

func main() {
	r := getRouter()
	http.ListenAndServe(":8080", r)
}
