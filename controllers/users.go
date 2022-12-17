package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/database"
  "github.com/quangminhvo79/go-api/authentication"
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
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user })
}

// POST /api/users/register
func CreateUser(c *gin.Context) {
  var input models.CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  user := models.User {
    Email: input.Email,
    Password: input.Password,
    Username: input.Username,
    AchievementWeightFrom: input.AchievementWeightFrom,
    AchievementWeightTo: input.AchievementWeightTo }

  if err := user.HashPassword(user.Password); err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  record := database.DB.Create(&user)
  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

// PATCH /api/users/:id
func UpdateUser(c *gin.Context) {
  var user models.User

  if err := authentication.UserScope.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "status": true, "error": "record not found" })
    return
  }

  var input models.UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": true, "error": err.Error() })
    return
  }

  if input.Password != "" {
    if err := user.HashPassword(input.Password); err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
  }

  updateUser := models.User {
    Email: input.Email,
    Password: user.Password,
    Username: input.Username,
    AchievementWeightFrom: input.AchievementWeightFrom,
    AchievementWeightTo: input.AchievementWeightTo }

  authentication.UserScope.Model(&user).Updates(&updateUser)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user.UserResponseData()})
}

// DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
  var user models.User
  if err := authentication.UserScope.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  authentication.UserScope.Delete(&user)

  c.JSON(http.StatusOK, gin.H{ "status": true })
}
