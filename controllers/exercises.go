package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
)

// GET /exercises
func FindExercises(c *gin.Context) {
	var exercises []models.Exercise
	database.DB.Find(&exercises)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": exercises })
}

// GET /exercises/:id_or_name
func FindExercisesBy(c *gin.Context) {
	var exercise models.Exercise

	name := "%" + c.Param("id_or_name") + "%"
	id := c.Param("id_or_name")

  if err := database.DB.Where("LOWER(name) LIKE ? OR id = ?", name, id).First(&exercise).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": exercise})
}

