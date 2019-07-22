package exercise

import (
	"context"
	"log"
	"net/http"

	"github.com/npthinhdev/valexer/internal/app/types"
	"github.com/npthinhdev/valexer/internal/pkg/response"
)

type (
	service interface {
		GetAll(ctx context.Context) ([]types.Exercise, error)
		Get(ctx context.Context, id string) (*types.Exercise, error)
		Create(ctx context.Context, exercise types.Exercise) error
		Update(ctx context.Context, id string, exercise types.Exercise) error
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
