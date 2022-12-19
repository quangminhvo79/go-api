package controllers

import (
	"net/http"
  "time"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/database"
  "github.com/quangminhvo79/go-api/scopes"
)

// GET /api/meal_histories
func FindMealHistories(c *gin.Context) {
	var meal_histories []models.MealHistory

  results := mealHistoriesScopes().Find(&meal_histories)
  if results.Error != nil || results.RowsAffected == 0 {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}

// GET /api/meal_histories/:id
func FindMealHistory(c *gin.Context) {
  var meal_history models.MealHistory

  results := mealHistoriesScopes().Where("meal_histories.id = ?", c.Param("id")).First(&meal_history)
  if results.Error != nil || results.RowsAffected == 0 {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": meal_history })
}

// POST /api/meal_histories
func CreateMealHistory(c *gin.Context) {
  var input models.MealHistoryInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var meal_history models.MealHistory
  meal_history.AssignAttributes(input)
  database.DB.Create(&meal_history)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": meal_history })
}

// PATCH /api/meal_histories/:id
func UpdateMealHistory(c *gin.Context) {
  var input models.MealHistoryInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var meal_history models.MealHistory
  meal_history.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&meal_history)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": meal_history })
}

// DELETE /api/meal_histories/:id
func DeleteMealHistory(c *gin.Context) {
  database.DB.Delete(&models.MealHistory{}, c.Param("id"))

  c.JSON(http.StatusOK, gin.H{ "status": true })
}

// GET /meal_histories/:user_id/achievement_rate
func AchievementRate(c *gin.Context) {
  // currently i do not have formula to calculation Achievement Rate
  // so i will temporary place achievement rate by 75%

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": gin.H{ "achievement_rate": 75, "date": time.Now() } })
}

// GET /meal_histories/:user_id/body_fat_percent_graph
func BodyFatPercentageGraph(c *gin.Context) {
  var meal_histories []models.BodyFatPercentGraph

  results := mealHistoriesScopes().Select("meal_histories.*", "SUM(Dish.calories) as total_calories").
                                   Group("DATE(meal_histories.date)").
                                   Find(&meal_histories);

  if results.Error != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "result": nil })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}

func mealHistoriesScopes() *gorm.DB {
  return database.DB.Scopes(scopes.MealHistories)
}
