package exercise

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/npthinhdev/valexer/internal/app/types"
	"github.com/npthinhdev/valexer/internal/pkg/response"
)

type (
	service interface {
		GetAll(ctx context.Context) ([]types.Exercise, error)
		Get(ctx context.Context, id string) (*types.Exercise, error)
		Create(ctx context.Context, exercise types.Exercise) (string, error)
		Update(ctx context.Context, exercise types.Exercise) error
		Delete(ctx context.Context, id string) error
	}
	// Handler is web handler
	Handler struct {
		srv service
	}
)

// New return new web handler
func New(s service) *Handler {
	return &Handler{s}
}

// GetAll handle get request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.srv.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, exercises)
}

// Get handle get exercise HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	exercise, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		log.Panicln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, exercise)
}

// Create handle insert exercise HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var exercise types.Exercise
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exercise.Title = r.Form.Get("title")
	exercise.Description = r.Form.Get("description")
	exercise.Testcase = r.Form.Get("testcase")

	id, err := h.srv.Create(r.Context(), exercise)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exercise.ID = id
	response.JSON(w, http.StatusCreated, nil)
}

// Update handle modify exercise HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var exercise types.Exercise
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exercise.ID = mux.Vars(r)["id"]
	exercise.Title = r.Form.Get("title")
	exercise.Description = r.Form.Get("description")
	exercise.Testcase = r.Form.Get("testcase")

	err = h.srv.Update(r.Context(), exercise)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}

// Delete handle delete exercise HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.srv.Delete(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, nil)
}
