package events

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_Create_InvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	RegisterHandler(r, NewService(&mockRepo{}))

	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/event", bytes.NewBuffer([]byte("{}")))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandler_Create_InvalidEventType(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	RegisterHandler(r, NewService(&mockRepo{}))

	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/event",
		bytes.NewBuffer([]byte(`{
			"user_id": "fcdf9dee-d759-44c9-a811-f82ae5000173",
			"type": "INVALID_EVENT_TYPE"
		}`)))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandler_Create_InvalidUserUUID(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	RegisterHandler(r, NewService(&mockRepo{}))

	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/event",
		bytes.NewBuffer([]byte(`{
			"user_id": " - must be uuid -",
			"type": "CLICK"
		}`)))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandler_Create_ValidRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	RegisterHandler(r, NewService(&mockRepo{}))

	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/event",
		bytes.NewBuffer([]byte(`{
			"user_id": "fcdf9dee-d759-44c9-a811-f82ae5000173",
			"type": "CLICK"
		}`)))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_Create_ESError(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	RegisterHandler(r, NewService(&mockRepo{errors.New("unit test")}))

	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/event",
		bytes.NewBuffer([]byte(`{
			"user_id": "fcdf9dee-d759-44c9-a811-f82ae5000173",
			"type": "CLICK"
		}`)))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
