package api

import (
	"net/http"

	"github.com/gorilla/mux"
	exerciseHandler "github.com/npthinhdev/valexer/internal/app/api/handler/exercise"
	"github.com/npthinhdev/valexer/internal/app/exercise"
)

type (
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
func Load(r *mux.Router) error {
	var exerciseRepo exercise.Repository
	exerciseSrv := exercise.NewService(exerciseRepo)
	exerciseAPI := exerciseHandler.New(exerciseSrv)

	routes := []route{
		// exercise
		{"/exercise", get, exerciseAPI.GetAll},
	}
	
	for _, rt := range routes {
		h := rt.handler
		r.HandleFunc(rt.path, h).Methods(rt.method)
	}
	return nil
}
