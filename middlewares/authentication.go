package middlewares

import (
	"github.com/quangminhvo79/go-api/authentication"
	"github.com/quangminhvo79/go-api/policies"
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

		loadPolicyScopes()

		c.Next()
	}
}

func loadPolicyScopes() {
	authentication.UserScope = policies.UserScope(authentication.Claims.Email)
}
