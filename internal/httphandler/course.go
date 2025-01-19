package httphandler

import (
	"ApiGateway/internal/models/courses"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// TODO: fix user id
	testUserID   = "user_2"
	testCourseID = "1"
)

func (h *Handler) createCourse(c *gin.Context) {
	var input courses.InputCourse

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
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

	req, err := http.NewRequest("POST", "http://localhost:8080/courses/create", bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
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

	if response.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
	}
}

func (h *Handler) getAllCourses(c *gin.Context) {
	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/courses/all", nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	defer resp.Body.Close()

	var response struct {
		Courses []courses.Courses `json:"courses"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, response.Courses)
	h.logg.Debug("Courses retrieved successfully")
}

func (h *Handler) getCourseByID(c *gin.Context) {
	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}
	courseID := c.Param("id")
	h.logg.Debug("DBG_courseID:", courseID)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/courses/id/"+courseID, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	defer resp.Body.Close()

	var course courses.Courses
	if err := json.NewDecoder(resp.Body).Decode(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, course)
	h.logg.Debug("Course retrieved successfully")
}

func (h *Handler) updateCourse(c *gin.Context) {
	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	courseID := c.Param("id")

	var input courses.UpdateCourse

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

	req, err := http.NewRequest("PUT", "http://localhost:8080/courses/update/"+courseID, bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
	}
	h.logg.Debug("Course updated successfully")
}

func (h *Handler) deleteCourse(c *gin.Context) {
	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	courseID := c.Param("id")

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:8080/courses/delete/"+courseID, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
	}
	h.logg.Debug("Course deleted successfully")
}

func (h *Handler) joinCourse(c *gin.Context) {
	token, err := h.services.GetToken(c.Request.Context(), testUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		h.logg.Error(err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8080/courses/join/"+testCourseID, nil)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		h.logg.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"message": "Join successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join course"})
	}
	h.logg.Debug("Join course successfully")
}
