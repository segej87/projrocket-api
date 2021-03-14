// controllers/people.go

package congrollers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
)

type CreatePersonInput struct {
	FirstName  string `json:"firstname" binding:"required"`
	LastName   string `json:"lastname" binding:"required"`
	Title      string `json:"title"`
}

// GET /people
// Get all people, or find by id
func FindPeople(c *gin.Context) {
  if id := c.Query("id"); id == "" {
    var people []models.Person
    models.DB.Find(&people)

    c.JSON(http.StatusOK, gin.H{"data": people})
  } else {
    var person models.Person

    if err := models.DB.Where("id = ?", c.Query("id")).First(&person).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"data": person})
  }
}

