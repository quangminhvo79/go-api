package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
  "github.com/quangminhvo79/go-api/database"
  "github.com/quangminhvo79/go-api/scopes"
  "github.com/quangminhvo79/go-api/authentication"
)

// GET /api/user_exercises
func FindUserExercises(c *gin.Context) {
	var user models.User
	userExerciseScope().Preload("Exercises").Find(&user)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": user.UserResponseData() })
}

// POST /api/user_exercises
func CreateUserExercises(c *gin.Context) {
	var userExercise models.UserExercise

  if err := c.ShouldBindJSON(&userExercise); err != nil {
  	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
  	return
  }

  userExercise.UserID = authentication.UserID
  record := database.DB.Create(&userExercise)
  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "error": record.Error.Error()})
    return
  }

  c.JSON(http.StatusCreated, gin.H{ "status": true, "exercises": userExercise })
}

// PATCH /api/user_exercises/:id
func UpdateUserExercises(c *gin.Context) {
	var input models.UserExerciseInput

  if err := c.ShouldBindJSON(&input); err != nil {
  	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
  	return
  }

  var userExercise models.UserExercise
  userExercise.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&userExercise)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": userExercise })
}

// DELETE /api/user_exercies/:id
func DeleteUserExercises(c *gin.Context) {
	database.DB.Delete(&models.UserExercise{}, c.Param("id"))
  c.JSON(http.StatusOK, gin.H{ "status": true })
}

func userExerciseScope() *gorm.DB {
	return database.DB.Scopes(scopes.UserExercise)
}
