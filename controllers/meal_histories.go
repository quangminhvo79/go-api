package controllers

import (
	"net/http"
  "time"
  "math"
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

// GET /api/achievement_rate
func AchievementRate(c *gin.Context) {
  var user models.User
  database.DB.Scopes(scopes.User).First(&user)


  caloriesTarget := (user.AchievementWeightTo - user.AchievementWeightFrom)

  if caloriesTarget < 0 {
    var achievement models.AchievementRate
    database.DB.Model(&models.ExerciseHistory{}).Where("diary_id = 1").Select("SUM(Exercise.calories_burned) as total_calories_burned").
              Joins("Diary").Joins("Exercise").Group("diary_id").
              First(&achievement)

    rate := (achievement.TotalCaloriesBurned / math.Abs(float64(caloriesTarget * 1000)) ) * 100
    c.JSON(http.StatusOK, gin.H{ "status": true, "result": gin.H{ "achievement_rate": math.Round(rate*100), "date": time.Now() } })
  } else {
    var meal_history models.BodyFatPercentGraph
    mealHistoriesScopes().Select("meal_histories.*", "SUM(Dish.calories) as total_calories").
                                   Group("meal_histories.user_id").
                                   First(&meal_history);

    rate := (meal_history.TotalCalories / math.Abs(float64(caloriesTarget * 10000)) ) * 100
    c.JSON(http.StatusOK, gin.H{ "status": true, "result": gin.H{ "achievement_rate": math.Round(rate*100), "date": time.Now() } })
  }
}

// GET /api/body_fat_percent_graph
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
