package events

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// Repository encapsulates data access logic from and to ES.
type Repository interface {
	Create(ctx context.Context, event *Event) error
}

type repository struct {
	es *elastic.Client
}

// NewRepository creates a new event repository
func NewRepository(es *elastic.Client) Repository {
	return &repository{es}
}

// Create persists the event into ES
func (r *repository) Create(ctx context.Context, event *Event) error {
	date := event.CreatedAt.Format("2006_01_02")
	_, err := r.es.Index().
		Index(fmt.Sprintf("events_%s", date)).
		BodyJson(event).
		Do(ctx)
	return err
}
