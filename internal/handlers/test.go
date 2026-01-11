package handlers

import (
	"credilens-backend/internal/models"
	"credilens-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
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
