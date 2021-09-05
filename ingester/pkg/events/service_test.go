package events

import (
	"context"
	"testing"
)

func TestService_Create(t *testing.T) {
	s := NewService(&mockRepo{})

	req := &CreateEventRequest{
		UserID: "fcdf9dee-d759-44c9-a811-f82ae5000173",
		Type:   ClickEvent,
	}

	err := s.Create(context.Background(), req)

	if err != nil {
		t.Errorf("Expected err to be nil, but got %s", err)
	}
}
