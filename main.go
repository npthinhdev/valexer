package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/npthinhdev/valexer/internal/pkg/db"

	"github.com/npthinhdev/valexer/internal/app/api"
	"github.com/npthinhdev/valexer/internal/app/router"
	"github.com/npthinhdev/valexer/internal/pkg/config/env"
	"github.com/npthinhdev/valexer/internal/pkg/db/mongodb"
)

type srvConfig struct {
	DB struct {
		Type    string `env:"DB_TYPE" default:"mongodb"`
		MongoDB mongodb.Config
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
	env.Load(&conf)

	conns := initInfraConns(&conf)
	defer conns.Close()

	log.Println("initializing HTTP routing...")
	getRouter, err := router.Init(conns)
	if err != nil {
		log.Panic("failed to init routing, err: ", err)
	}

	getAddr := fmt.Sprintf("%s:%d", conf.HTTP.Address, conf.HTTP.Port)
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

func initInfraConns(conf *srvConfig) *api.InfraConns {
	conns := &api.InfraConns{}
	conns.Database.Type = conf.DB.Type

	switch conf.DB.Type {
	case db.TypeMongoDB:
		s, err := mongodb.Dial(&conf.DB.MongoDB)
		if err != nil {
			log.Panic("failed to dial to target server, err: ", err)
		}
		conns.Database.MongoDB = s
	}
	return conns
}
