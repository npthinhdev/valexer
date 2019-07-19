package main

import (
	"net/http"

	"github.com/npthinhdev/valexer/internal/app"
)

func main() {
	r := app.NewRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
