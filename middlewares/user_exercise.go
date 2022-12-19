package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasUserExercise() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userExercise models.UserExercise
	  if err := database.DB.Scopes(scopes.UserExercise).Where("user_exercises.id = ?", c.Param("id")).First(&userExercise).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
	    return
	  }

	  c.Next()
	}
}
