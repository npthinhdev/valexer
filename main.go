package main

import (
	"net/http"
)

// Declare templates
var tmlp = getTemplates()

// Context hold value in template
type Context map[string]string

func main() {
	r := getRouter()
	http.ListenAndServe(":8080", r)
}
