package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/npthinhdev/valexer/internal/app/router"
)

type srvConfig struct {
	DB struct {
		Type string `env:"DB_TYPE" default:"mongodb"`
	}
	HTTP struct {
		Address           string        `env:"HTTP_ADDRESS" default:""`
		Port              int           `env:"PORT" default:"8080"`
		ReadTimout        time.Duration `env:"HTTP_READ_TIMEOUT" default:"5m"`
		WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"5m"`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" default:"30s"`
	}
}

func main() {
	var conf srvConfig
	// env.Load(&conf)

	conf.HTTP.Address = ""
	conf.HTTP.Port = 8080
	getAddr := fmt.Sprintf("%s:%d", conf.HTTP.Address, conf.HTTP.Port)

	log.Println("initializing HTTP routing...")
	getRouter, err := router.Init()
	if err != nil {
		log.Panic("failed to init routing, err:", err)
	}

	httpServer := http.Server{
		Addr:              getAddr,
		Handler:           getRouter,
		ReadTimeout:       conf.HTTP.ReadTimout,
		WriteTimeout:      conf.HTTP.WriteTimeout,
		ReadHeaderTimeout: conf.HTTP.ReadHeaderTimeout,
	}

	log.Println("starting HTTP server...")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Panic("http.ListenAndServe() error:", err)
	}
}
