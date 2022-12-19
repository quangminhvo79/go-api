package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func HasPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
	  if err := database.DB.Scopes(scopes.Post).Where("posts.id = ?", c.Param("id")).First(&post).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
	    return
	  }

	  c.Next()
	}
}
