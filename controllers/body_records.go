package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
)

// GET /api/body_records
func FindBodyRecords(c *gin.Context) {
	var bodyRecords []models.BodyRecord

	err := bodyRecordScope().Find(&bodyRecords).Error
	if err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": bodyRecords })
}

// POST /api/body_records
func CreateBodyRecord(c *gin.Context) {
	var input models.BodyRecordInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }
  var bodyRecord models.BodyRecord
  bodyRecord.AssignAttributes(input)
  record := database.DB.Create(&bodyRecord)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusCreated, gin.H{ "status": true, "result": bodyRecord })
}

// PATCH /api/body_records/:id
func UpdateBodyRecord(c *gin.Context) {
  var input models.BodyRecordInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var bodyRecord models.BodyRecord
  bodyRecord.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&bodyRecord)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": bodyRecord })
}

func DeleteBodyRecord(c *gin.Context) {
  database.DB.Delete(&models.BodyRecord{}, c.Param("id"))

  c.JSON(http.StatusOK, gin.H{ "status": true })
}

func bodyRecordScope() *gorm.DB {
	return database.DB.Scopes(scopes.BodyRecord)
}
