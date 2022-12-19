package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
  "gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/database"
  "github.com/quangminhvo79/go-api/scopes"
)

// GET /api/users
func FindUsers(c *gin.Context) {
	var users []models.UserOutput
	database.DB.Model(&models.User{}).Find(&users)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": users })
}

// GET /api/users/:id
func FindUser(c *gin.Context) {
  var user models.UserOutput

  if err := database.DB.Model(&models.User{}).Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user })
}

// POST /api/users/register
func CreateUser(c *gin.Context) {
  var input models.UserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var user models.User
  user.AssignAttributes(input)

  if err := user.HashPassword(user.Password); err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": err.Error() })
    return
  }

  database.DB.Create(&user)
  c.JSON(http.StatusCreated, gin.H{ "status": true, "user": gin.H{ "user_id": user.ID, "email": user.Email, "username": user.Username }} )
}

// PATCH /api/users/:id
func UpdateUser(c *gin.Context) {
  var input models.UserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": true, "error": err.Error() })
    return
  }

  var user models.User
  user.AssignAttributes(input)

  if input.Password != "" {
    if err := user.HashPassword(input.Password); err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": err.Error() })
      return
    }
  }

  userScopes().Updates(&user)
  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user.UserResponseData()})
}

// DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
  userScopes().Delete(c.Param("id"))
  c.JSON(http.StatusOK, gin.H{ "status": true })
}

func userScopes() *gorm.DB {
  return database.DB.Scopes(scopes.User)
}
