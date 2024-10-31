package httphandler

import (
	"ApiGateway/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.InputSignUp

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	body, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8081/users/sign-up", bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	var response struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"id": response.ID,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.InputSignIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	body, err := json.Marshal(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8081/users/sign-in", bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	var response struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	h.logg.Info("RESPONSE", slog.String("response", fmt.Sprintf("%+v", response.ID)))

	token, err := h.grpcClient.GetToken(c.Request.Context(), int64(response.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	redisID := string(rune(response.ID))

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
