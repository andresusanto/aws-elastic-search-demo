package events

import (
	"testing"
)

func TestCreateEventRequest_Validate_Success(t *testing.T) {
	valids := []EventType{ClickEvent}

	for _, v := range valids {
		req := &CreateEventRequest{
			Type: v,
		}

		err := req.Validate()

		if err != nil {
			t.Errorf("Expected err to be nil, but got %s", err.Error())
		}
	}
}

func TestCreateEventRequest_Validate_Error(t *testing.T) {
	invalids := []EventType{"11", "22", "33"}
	expect := "bad event type"

	for _, i := range invalids {
		req := &CreateEventRequest{
			Type: i,
		}

		err := req.Validate()

		if err == nil || err.Error() != expect {
			t.Errorf("Expected err %s, but got %s", expect, err)
		}
	}
}
