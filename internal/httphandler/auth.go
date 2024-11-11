package httphandler

import (
	"ApiGateway/internal/models/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input auth.InputSignUp

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	id, err := h.userClient.SignUp(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input auth.InputSignIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	id, err := h.userClient.SignIn(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	h.logg.Debug("RESPONSE", slog.String("response", fmt.Sprintf("%+v", id)))

	token, err := h.grpcClient.GetToken(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	redisID := fmt.Sprintf("user_%d", id)

	err = h.services.SetToken(c.Request.Context(), redisID, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set token in Redis"})
		h.logg.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
