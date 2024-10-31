package httphandler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.logg.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
