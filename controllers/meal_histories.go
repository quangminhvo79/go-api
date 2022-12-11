package controllers

import (
	"net/http"
  "strconv"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/quangminhvo79/go-api/models"
)

// GET /meal_histories/:user_id
func FindMealHistories(c *gin.Context) {
	var meal_histories []models.MealHistory

  uID, _ := strconv.Atoi(c.Param("user_id"))
  results := models.MealHistories().Scopes(models.MealHistoriesByUser(uID)).Find(&meal_histories);

  if results.Error != nil || results.RowsAffected == 0 {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}

// GET /meal_histories/:user_id/session/:session_id
func FindMealHistoriesBySession(c *gin.Context) {
  var meal_histories []models.MealHistory

  sID, _ := strconv.Atoi(c.Param("session_id"))
  uID, _ := strconv.Atoi(c.Param("user_id"))

  results := models.MealHistories().Scopes(
    models.MealHistoriesByUser(uID),
    models.MealHistoriesBySession(sID),
  ).Find(&meal_histories);

  if results.Error != nil || results.RowsAffected == 0  {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}

// POST /meal_histories
func CreateMealHistory(c *gin.Context) {
  var input models.CreateMealHistoryInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  meal_history := models.MealHistory{
    Date: input.Date,
    SessionID: input.SessionID,
    DishID: input.DishID,
    UserID: input.UserID,
  }

  models.DB.Create(&meal_history)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": meal_history })
}

// PATCH /meal_histories/:id
func UpdateMealHistory(c *gin.Context) {
  var meal_history models.MealHistory
  ID, _ := strconv.Atoi(c.Param("id"))

  if err := models.MealHistoryByID(ID).First(&meal_history).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "status": false, "error": "Record not found" })
    return
  }

  var input models.UpdateMealHistoryInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  updateMealHistory := models.MealHistory{
    Date: input.Date,
    SessionID: input.SessionID,
    DishID: input.DishID,
    UserID: input.UserID}

  models.DB.Model(&meal_history).Updates(&updateMealHistory)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": updateMealHistory })
}

// DELETE /meal_histories/:id
func DeleteMealHistory(c *gin.Context) {
  var meal_history models.MealHistory
  ID, _ := strconv.Atoi(c.Param("id"))

  if err := models.MealHistoryByID(ID).First(&meal_history).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  models.DB.Delete(&meal_history)

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
  uID, _ := strconv.Atoi(c.Param("user_id"))

  results := models.MealHistories().Scopes(models.MealHistoriesByUser(uID)).
    Select("meal_histories.*", "SUM(Dish.calories) as total_calories").
    Group("DATE(meal_histories.date)").
    Find(&meal_histories);

  if results.Error != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "total": results.RowsAffected, "result": meal_histories })
}
