package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	  var user models.User

	  if err := database.DB.Scopes(scopes.User).Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "status": true, "error": "record not found" })
	    return
	  }

	  c.Next()
	}
}
