// controllers/people.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  //"fmt"
  //"strings"
)

type CreatePersonInput struct {
	FirstName  string `form:"firstname" json:"firstname" binding:"required"`
	LastName   string `form:"lastname" json:"lastname" binding:"required"`
	Title      string `form:"title" json:"title"`
}

type SearchPersonInput struct {
	FirstName  string `form:"firstname" json:"firstname"`
	LastName   string `form:"lastname" json:"lastname"`
	Title      string `form:"title" json:"title"`
}

// GET /people
// Get all people, or find by id
func FindPeople(c *gin.Context) {
  query := c.Request.URL.Query()

  if len(query) == 0 {
    var people []models.Person
    models.DB.Find(&people)

    c.JSON(http.StatusOK, gin.H{"data": people})
  } else {
    var person models.Person
    var searchPerson SearchPersonInput

    if c.BindQuery(&searchPerson) == nil {
      if err := models.DB.Where(&models.Person{FirstName: searchPerson.FirstName, LastName: searchPerson.LastName, Title: searchPerson.Title}).First(&person).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }
    }
    
    c.JSON(http.StatusOK, gin.H{"data": person})
  }
}

// POST /people
// Create new person
func CreatePerson(c *gin.Context) {
  // Validate input
  var input CreatePersonInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create person
  person := models.Person{FirstName: input.FirstName, LastName: input.LastName, Title: input.Title}
  models.DB.Create(&person)

  c.JSON(http.StatusOK, gin.H{"data": person})
}

