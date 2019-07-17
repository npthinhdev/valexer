package main

import (
	"net/http"

	"github.com/npthinhdev/valexer/app"
)

func main() {
	r := app.GetRouter()
	http.ListenAndServe(":8080", r)
}
