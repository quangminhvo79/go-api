package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasBodyRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyRecord models.BodyRecord
	  if err := database.DB.Scopes(scopes.BodyRecord).Where("body_records.id = ?", c.Param("id")).First(&bodyRecord).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
	    return
	  }

	  c.Next()
	}
}
