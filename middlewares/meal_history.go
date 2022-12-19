package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasMealHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
  	var meal_history models.MealHistory
	  if err := database.DB.Scopes(scopes.MealHistories).Where("meal_histories.id = ?", c.Param("id")).First(&meal_history).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "status": false, "error": "Record not found" })
	    return
	  }

	  c.Next()
	}
}
