package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/authentication"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	tokenString, err:= authentication.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
