package api

import (
	"net/http"

	"github.com/gorilla/mux"
	exerciseHandler "github.com/npthinhdev/valexer/internal/app/api/handler/exercise"
	"github.com/npthinhdev/valexer/internal/app/exercise"
	"github.com/npthinhdev/valexer/internal/pkg/db"
)

type (
	// InfraConns holds infrastructure services connections like MongoDB, Redis, Kafka,...
	InfraConns struct {
		Database db.Connections
	}
	route struct {
		path    string
		method  string
		handler http.HandlerFunc
	}
)

const (
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	delete = http.MethodDelete
)

// Load is transfer api path to route
func Load(r *mux.Router, conns *InfraConns) error {
	var exerciseRepo exercise.Repository

	switch conns.Database.Type {
	case db.TypeMongoDB:
		exerciseRepo = exercise.NewMongoRepository(conns.Database.MongoDB)
	}

	exerciseSrv := exercise.NewService(exerciseRepo)
	exerciseAPI := exerciseHandler.New(exerciseSrv)

	routes := []route{
		// exercise
		{"/exercise", get, exerciseAPI.GetAll},
		{"/exercise/{id}", get, exerciseAPI.Get},
		{"/exercise", post, exerciseAPI.Create},
		{"/exercise/{id}", put, exerciseAPI.Update},
		{"/exercise/{id}", delete, exerciseAPI.Delete},
	}

	for _, rt := range routes {
		h := rt.handler
		r.HandleFunc("/api"+rt.path, h).Methods(rt.method)
	}
	return nil
}

// Close close all underlying connections
func (c *InfraConns) Close() {
	c.Database.Close()
}
