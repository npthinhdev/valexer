package main

import (
	"net/http"
	"valexer/app"
)

func main() {
	r := app.GetRouter()
	http.ListenAndServe(":8080", r)
}
