package events

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// RegisterHandler sets up the routing of the HTTP handlers.
func RegisterHandler(r *gin.Engine, s Service) {
	h := &handler{s}
	r.Group("/v1").
		POST("/event", h.create)
}

type handler struct {
	service Service
}

func (h *handler) create(c *gin.Context) {
	req := CreateEventRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(c.Request.Context(), req); err != nil {
		reqID := requestid.Get(c)
		log.Error().Err(err).Str("reqId", reqID).Msg("unknown error occured when creating event")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
