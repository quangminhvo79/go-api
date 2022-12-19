package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/quangminhvo79/go-api/models"
	"github.com/quangminhvo79/go-api/database"
	"github.com/quangminhvo79/go-api/scopes"
)

// GET /api/posts
func FindPosts(c *gin.Context) {
	var posts []models.Post
  database.DB.Model(&models.Post{}).Find(&posts)

	c.JSON(http.StatusOK, gin.H{ "status": true, "result": posts })
}

// GET /api/posts
func FindPost(c *gin.Context) {
  var post models.Post

  if err := database.DB.Model(&models.Post{}).Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": "Record not found!" })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": post })
}

// POST /api/posts
func CreatePost(c *gin.Context) {
	var input models.PostInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }
  var post models.Post
  post.AssignAttributes(input)
  record := database.DB.Create(&post)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusCreated, gin.H{ "status": true, "result": post })
}

// PATCH /api/posts/:id
func UpdatePost(c *gin.Context) {
  var input models.PostInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "status": false, "error": err.Error() })
    return
  }

  var post models.Post
  post.AssignAttributes(input)
  record := database.DB.Where("id = ?", c.Param("id")).Updates(&post)

  if record.Error != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "status": false, "error": record.Error.Error() })
    return
  }

  c.JSON(http.StatusOK, gin.H{ "status": true, "result": post })
}

// DELETE /api/posts/id
func DeletePost(c *gin.Context) {
  database.DB.Delete(&models.Post{}, c.Param("id"))
  c.JSON(http.StatusOK, gin.H{ "status": true })
}

func postScope() *gorm.DB {
	return database.DB.Scopes(scopes.Post)
}

