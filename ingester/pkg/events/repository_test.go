package events

import (
	"context"
	"errors"
	"time"
)

type mockRepo struct {
	err error
}

func (r *mockRepo) Create(ctx context.Context, event *Event) error {
	if event.Type == "" {
		return errors.New("empty type")
	}
	if event.UserID == "" {
		return errors.New("empty user id")
	}
	if time.Since(event.CreatedAt) > 1*time.Second {
		return errors.New("empty time")
	}
	return r.err
}
