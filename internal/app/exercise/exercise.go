package exercise

import (
	"context"

	"github.com/npthinhdev/valexer/internal/app/types"
)

// Repository is an interface of a exercise repository
type Repository interface {
	FindAll(ctx context.Context) ([]types.Exercise, error)
	FindByID(ctx context.Context, id string) (*types.Exercise, error)
	Create(ctx context.Context, exercise types.Exercise) (string, error)
	Update(ctx context.Context, exercise types.Exercise) error
	Delete(ctx context.Context, id string) error
}

// Service is an exercise service
type Service struct {
	repo Repository
}

// NewService return a new exercise service
func NewService(r Repository) *Service {
	return &Service{r}
}

// GetAll return all exercise from database
func (s *Service) GetAll(ctx context.Context) ([]types.Exercise, error) {
	return s.repo.FindAll(ctx)
}

// Get return given exercise by id
func (s *Service) Get(ctx context.Context, id string) (*types.Exercise, error) {
	return s.repo.FindByID(ctx, id)
}

// Create a exercise
func (s *Service) Create(ctx context.Context, excercise types.Exercise) (string, error) {
	return s.repo.Create(ctx, excercise)
}

// Update a exercise
func (s *Service) Update(ctx context.Context, exercise types.Exercise) error {
	_, err := s.repo.FindByID(ctx, exercise.ID)
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, exercise)
}

// Delete a exercise
func (s *Service) Delete(ctx context.Context, id string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
