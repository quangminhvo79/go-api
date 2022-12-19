package middlewares

import (
	"net/http"
	"github.com/quangminhvo79/go-api/authentication"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "request does not contain an access token"})
			return
		}

		err:= authentication.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := database.DB.Scopes(scopes.User).First(&user).Error; err != nil {
	    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Invalid credentials!" })
	    return
	  }
	  authentication.UserID = user.ID

		c.Next()
	}
}
