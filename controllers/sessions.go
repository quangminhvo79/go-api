package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
)

// GET /sessions
func FindSessions(c *gin.Context) {
	var sessions []models.Session
	database.DB.Find(&sessions)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": sessions })
}

// GET /sessions/:id
func FindSessionsBy(c *gin.Context) {
	var session models.Session

	name := "%" + c.Param("id") + "%"
	id := c.Param("id")

  if err := database.DB.Where("LOWER(name) LIKE ? OR id = ?", name, id).First(&session).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false,"error": "Record not found!" })
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": session })
}

// GET /sessions/:id/meal_histories
func FindMealHistoriesBySession(c *gin.Context) {
  var meal_histories []models.MealHistory

  results := database.DB.Scopes(scopes.MealHistories).Where("session_id = ?", c.Param("id")).Find(&meal_histories)
  if results.Error != nil || results.RowsAffected == 0  {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}
