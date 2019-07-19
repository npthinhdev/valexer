package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/npthinhdev/valexer/internal/app"
)

type srvConfig struct {
	DB struct {
		Type string
	}
	HTTP struct {
		Address           string
		Port              int
		ReadTimout        time.Duration
		WriteTimeout      time.Duration
		ReadHeaderTimeout time.Duration
	}
}

func main() {
	var conf srvConfig

	conf.HTTP.Address = "127.0.0.1"
	conf.HTTP.Port = 8080
	addr := fmt.Sprintf("%s:%d", conf.HTTP.Address, conf.HTTP.Port)
	router := app.NewRouter()

	httpServer := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       conf.HTTP.ReadTimout,
		WriteTimeout:      conf.HTTP.WriteTimeout,
		ReadHeaderTimeout: conf.HTTP.ReadHeaderTimeout,
	}

	log.Println("starting HTTP server...")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Panic("http.ListenAndServe() error:", err)
	}
}
