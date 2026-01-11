package handlers

import (
	"net/http"

	"credilens-backend/internal/models"
	"credilens-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func Analyze(c *gin.Context) {
	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	result := services.AnalyzeContent(req.Content)
	c.JSON(http.StatusOK, result)
}
