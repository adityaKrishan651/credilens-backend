package ingest

import (
	"credilens-backend/internal/models"
	"credilens-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngestRequest struct {
	URL string `json:"url"`
}

func Ingest(c *gin.Context) {

	var req models.IngestRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	result, err := services.Scrape(req.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, result)

}
