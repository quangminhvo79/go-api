package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
)

// GET /users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": users })
}

// GET /users/:id
func FindUser(c *gin.Context) {
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user })
}

// POST /user
func CreateUser(c *gin.Context) {
  var input models.CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  user := models.User{
  	Email: input.Email,
  	Password: input.Password,
  	Name: input.Name,
  	AchievementWeightFrom: input.AchievementWeightFrom,
  	AchievementWeightTo: input.AchievementWeightTo }

  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": user })
}

// PATCH /users/:id
func UpdateUser(c *gin.Context) {
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "status": true, "error": "record not found" })
    return
  }

  var input models.UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": true, "error": err.Error() })
    return
  }

  updateUser := models.User{
  	Email: input.Email,
  	Password: input.Password,
  	Name: input.Name,
  	AchievementWeightFrom: input.AchievementWeightFrom,
  	AchievementWeightTo: input.AchievementWeightTo }

  models.DB.Model(&user).Updates(&updateUser)

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": updateUser })
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  models.DB.Delete(&user)

  c.JSON(http.StatusOK, gin.H{ "status": true,  })
}
