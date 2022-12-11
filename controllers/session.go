package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
)

// GET /sessions
func FindSessions(c *gin.Context) {
	var sessions []models.Session
	models.DB.Find(&sessions)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": sessions })
}

// GET /sessions/:name
func FindSessionsByName(c *gin.Context) {
	var session models.Session

  if err := models.DB.Where("LOWER(name) LIKE ?", "%" + c.Param("name") + "%").First(&session).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": session })
}
