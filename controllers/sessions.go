package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
)

// GET /sessions
func FindSessions(c *gin.Context) {
	var sessions []models.Session
	database.DB.Find(&sessions)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": sessions })
}

// GET /sessions/:id_or_name
func FindSessionsByName(c *gin.Context) {
	var session models.Session

	name := "%" + c.Param("id_or_name") + "%"
	id := c.Param("id_or_name")

  if err := database.DB.Where("LOWER(name) LIKE ? OR id = ?", name, id).First(&session).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": session })
}
