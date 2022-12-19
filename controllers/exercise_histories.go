package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
)

// GET /api/diaries/:id/exercise_histories
func FindExerciseHistories(c *gin.Context) {
	var exercise_histories []models.ExerciseHistory
	database.DB.Where("diary_id = ?", c.Param("id")).Preload("Exercise").Find(&exercise_histories)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": exercise_histories })
}

// GET /api/exercise_history/:id
func FindExerciseHistory(c *gin.Context) {
  var exercise_history models.ExerciseHistory

  if err := database.DB.Where("id = ?", c.Param("id")).Preload("Exercise").First(&exercise_history).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": exercise_history })
}

// POST /api/exercise_histories
func CreateExerciseHistory(c *gin.Context) {
	var input models.ExerciseHistoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }
  var exercise_history models.ExerciseHistory
  exercise_history.AssignAttributes(input)
  record := database.DB.Create(&exercise_history)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusCreated, gin.H{ "status": true, "result": exercise_history })
}

// PATCH /api/exercise_histories/:id
func UpdateExerciseHistory(c *gin.Context) {
  var input models.ExerciseHistoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var exercise_history models.ExerciseHistory
  exercise_history.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&exercise_history)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": exercise_history })
}

// DELETE /api/exercise_histories/id
func DeleteExerciseHistory(c *gin.Context) {
  database.DB.Delete(&models.ExerciseHistory{}, c.Param("id"))
  c.JSON(http.StatusOK, gin.H{ "status": true })
}
