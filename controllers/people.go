// controllers/people.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "time"
  "github.com/satori/go.uuid"
  //"fmt"
  //"strings"
)

type CreatePersonInput struct {
  CreatedBy  uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  FirstName  string    `form:"firstname" json:"firstname" binding:"required"`
  LastName   string    `form:"lastname" json:"lastname" binding:"required"`
  Email      string    `form:"email" json:"email"`
  Phone      string    `form:"phone" json:"phone"`
  Birthday   time.Time `form:"birthday" json:"birthday"`
  Title      string    `form:"title" json:"title"`
  Department string    `form:"department" json:"department"`
  Self       string    `form:"string" json:"string" gorm:"default:false"`
}

type UpdatePersonInput struct {
  CreatedBy  uuid.UUID `form:"created_by" json:"created_by"`
  FirstName  string    `form:"firstname" json:"firstname"`
  LastName   string    `form:"lastname" json:"lastname"`
  Email      string    `form:"email" json:"email"`
  Phone      string    `form:"phone" json:"phone"`
  Birthday   time.Time `form:"birthday" json:"birthday"`
  Title      string    `form:"title" json:"title"`
  Department string    `form:"department" json:"department"`
  Self       string    `form:"string" json:"string"`
}

type SearchPersonInput struct {
  FirstName  string    `form:"firstname" json:"firstname"`
  LastName   string    `form:"lastname" json:"lastname"`
  Email      string    `form:"email" json:"email"`
  Phone      string    `form:"phone" json:"phone"`
  Birthday   time.Time `form:"birthday" json:"birthday"`
  Title      string    `form:"title" json:"title"`
  Department string    `form:"department" json:"department"`
  Self       bool      `form:"self" json:"self"`
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
  person := models.Person{CreatedBy: input.CreatedBy, FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, Phone: input.Phone, Birthday: input.Birthday, Title: input.Title, Department: input.Department}
  models.DB.Create(&person)

  c.JSON(http.StatusOK, gin.H{"data": person})
}

// GET /people
// Get all people, or find by id or query params
func FindPeople(c *gin.Context) {
  query := c.Request.URL.Query()

  var people []models.Person

  if len(query) == 0 {
    models.DB.Find(&people)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&people, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else if query.Get("created_by") != "" {
    var searchPerson SearchPersonInput

    if bindErr := c.BindQuery(&searchPerson); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Person{FirstName: searchPerson.FirstName, LastName: searchPerson.LastName, Email: searchPerson.Email, Phone: searchPerson.Phone, Birthday: searchPerson.Birthday, Title: searchPerson.Title, Department: searchPerson.Department, Self: searchPerson.Self}).Find(&people, "created_by = ?", query.Get("created_by")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchPerson SearchPersonInput

    if bindErr := c.BindQuery(&searchPerson); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Person{FirstName: searchPerson.FirstName, LastName: searchPerson.LastName, Email: searchPerson.Email, Phone: searchPerson.Phone, Birthday: searchPerson.Birthday, Title: searchPerson.Title, Self: searchPerson.Self}).Find(&people).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": people})
}

// PATCH /people/:id
// Update a person
func UpdatePerson(c *gin.Context) {
  // Get the person to be updated
  var person models.Person
  if err := models.DB.First(&person, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdatePersonInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&person).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": person})
}

// DELETE /people/:id
// Delete a person
func DeletePerson(c *gin.Context) {
  // Get model if exist
  var person models.Person
  if err := models.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&person)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
