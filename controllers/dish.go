package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
)

// GET /dishes
func FindDishes(c *gin.Context) {
	var dishes []models.Dish
	models.DB.Find(&dishes)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": dishes })
}

// GET /dishes/:name
func FindDishesByName(c *gin.Context) {
	var dish models.Dish

  if err := models.DB.Where("LOWER(name) LIKE ?", "%" + c.Param("name") + "%").First(&dish).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": dish})
}
