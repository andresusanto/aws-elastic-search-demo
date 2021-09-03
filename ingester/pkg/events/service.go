package events

import (
	"context"
	"time"
)

// Service encapsulates usecase logic for events
type Service interface {
	Create(ctx context.Context, req CreateEventRequest) error
}

type service struct {
	repo Repository
}

// NewService creates a new event service.
func NewService(r Repository) Service {
	return &service{r}
}

// Create creates a new event.
func (s *service) Create(ctx context.Context, req CreateEventRequest) error {
	event := &Event{
		req.UserID,
		req.Type,
		time.Now(),
	}

	return s.repo.Create(ctx, event)
}
