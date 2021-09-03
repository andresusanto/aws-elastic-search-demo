package events

import (
	"errors"
	"time"
)

// Valid event types
const (
	ClickEvent EventType = "CLICK"
)

// EventType represents type of event performed
type EventType string

// Event represents an action performed by user
type Event struct {
	UserID    string    `json:"user_id"`
	Type      EventType `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateEventRequest represents an incoming event
type CreateEventRequest struct {
	UserID string    `json:"user_id" binding:"required,uuid4"`
	Type   EventType `json:"type" binding:"required"`
}

// Validate validates the CreateEventRequest fields.
func (m *CreateEventRequest) Validate() error {
	if m.Type != ClickEvent {
		return errors.New("bad event type")
	}
	return nil
}
