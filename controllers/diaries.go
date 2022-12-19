package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
)

// GET /api/diaries
func FindDiaries(c *gin.Context) {
	var diaries []models.Diary
	diaryScope().Find(&diaries)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": diaries })
}

// GET /api/diaries/:id
func FindDiary(c *gin.Context) {
	var diary models.Diary

	if err := diaryScope().Where("id = ?", c.Param("id")).First(&diary).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": diary })
}

// POST /api/diaries
func CreateDiary(c *gin.Context) {
	var input models.DiaryInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }
  var diary models.Diary
  diary.AssignAttributes(input)
  record := database.DB.Create(&diary)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusCreated, gin.H{ "status": true, "result": diary })
}

// PATCH /api/diaries/:id
func UpdateDiary(c *gin.Context) {
  var input models.DiaryInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var diary models.Diary
  diary.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&diary)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": diary })
}

// DELETE /api/diaries/id
func DeleteDiary(c *gin.Context) {
  database.DB.Delete(&models.Diary{}, c.Param("id"))
  c.JSON(http.StatusOK, gin.H{ "status": true })
}

func diaryScope() *gorm.DB {
	return database.DB.Scopes(scopes.Diary)
}

