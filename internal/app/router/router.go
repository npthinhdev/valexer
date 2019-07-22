package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/npthinhdev/valexer/internal/app/api"
	"github.com/npthinhdev/valexer/internal/app/router/handler/admin"
	"github.com/npthinhdev/valexer/internal/app/router/handler/exercise"
	"github.com/npthinhdev/valexer/internal/app/router/handler/index"
)

type (
	route struct {
		path    string
		method  string
		handler http.HandlerFunc
	}
	static struct {
		prefix string
		dir    string
	}
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

// Init all router
func Init() (http.Handler, error) {
	indexView := index.New()
	adminView := admin.New()
	exerciseView := exercise.New()

	routes := []route{
		// home
		{"/", get, indexView.Index},
		// admin
		{"/admin", get, adminView.Index},
		// exercise
		{"/exercise/{id}", get, exerciseView.Get},
		{"/exercise/create/", get, exerciseView.CreateGET},
		{"/exercise/create/", post, exerciseView.CreatePOST},
		{"/exercise/edit/{id}", get, exerciseView.UpdateGET},
		{"/exercise/edit/{id}", post, exerciseView.UpdatePOST},
		{"/exercise/delete/{id}", get, exerciseView.Delete},
	}

	r := mux.NewRouter()
	err := api.Load(r)
	if err != nil {
		log.Println(err)
	}
	for _, rt := range routes {
		h := rt.handler
		r.HandleFunc(rt.path, h).Methods(rt.method)
	}

	s := static{
		prefix: "/static/",
		dir:    "web/static",
	}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}
