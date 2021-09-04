package healthcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler performs graceful shutdown and processes health check requests
type Handler interface {
	Shutdown()
}

// RegisterHandler sets up the routing of the HTTP handlers.
func RegisterHandler(r *gin.Engine) Handler {
	h := &handler{true}
	r.GET("/health", h.get)
	return h
}

type handler struct {
	ok bool
}

func (h *handler) Shutdown() {
	h.ok = false
	time.Sleep(10 * time.Second)
}

func (h *handler) get(c *gin.Context) {
	if h.ok {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusServiceUnavailable)
}
