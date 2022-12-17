package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
)

// GET /dishes
func FindDishes(c *gin.Context) {
	var dishes []models.Dish
	database.DB.Find(&dishes)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": dishes })
}

// GET /dishes/:id_or_name
func FindDishesByName(c *gin.Context) {
	var dish models.Dish

	name := "%" + c.Param("id_or_name") + "%"
	id := c.Param("id_or_name")

  if err := database.DB.Where("LOWER(name) LIKE ? OR id = ?", name, id).First(&dish).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": dish})
}
