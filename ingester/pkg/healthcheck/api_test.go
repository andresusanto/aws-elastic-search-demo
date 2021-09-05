package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_Shutdown(t *testing.T) {
	h := &handler{true, 0}

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/", h.get)

	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	h.Shutdown() // mock SIGTERM

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("Expected status %d, got %d", http.StatusServiceUnavailable, w.Code)
	}
}
