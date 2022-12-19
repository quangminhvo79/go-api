package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		var diary models.Diary
	  if err := database.DB.Scopes(scopes.Diary).Where("diaries.id = ?", c.Param("id")).First(&diary).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
	    return
	  }

	  c.Next()
	}
}
