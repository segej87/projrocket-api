// controllers/users.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  //"time"
  //"fmt"
  //"strings"
)

type CreateUserInput struct {
  Username  string `form:"username" json:"username" binding:"required"`
}

type SearchUserInput struct {
  Username  string `form:"username" json:"username"`
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
  // Validate input
  var input CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create user
  user := models.User{Username: input.Username}
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users
// Get all users, or find by id or username
func FindUsers(c *gin.Context) {
  query := c.Request.URL.Query()
  
  var users []models.User

  if len(query) == 0 {
    models.DB.Find(&users)
  } else if query.Get("id") != "" {
    if err := models.DB.Where("id = ?", query.Get("id")).Find(&users).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchUser SearchUserInput

    if c.BindQuery(&searchUser) == nil {
      if err := models.DB.Where(&models.User{Username: searchUser.Username}).Find(&users).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }
    }
  }
  
  c.JSON(http.StatusOK, gin.H{"data": users})
}

