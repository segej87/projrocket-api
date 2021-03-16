// controllers/relationships.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "github.com/satori/go.uuid"
)

type CreateRelationshipInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" binding:"required"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" binding:"required"`
  Type      string    `form:"type" json:"type"`
}

type UpdateRelationshipInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  FromID    uuid.UUID `form:"from_id" json:"from_id"`
  ToID      uuid.UUID `form:"to_id" json:"to_id"`
  Type      string    `form:"type" json:"type"`
}

type SearchRelationshipInput struct {
  Type string `form:"type" json:"type"`
}

// GET /relationships
// Get all relationships, or find by id or query params
func FindRelationships(c *gin.Context) {
  query := c.Request.URL.Query()

  var relationships []models.Relationship

  if len(query) == 0 {
    models.DB.Find(&relationships)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&relationships, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchRelationship SearchRelationshipInput

    // Validate input
    if bindErr := c.BindQuery(&searchRelationship); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    // TODO: try to fix this crazy block
    if query.Get("created_by") != "" && query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships, "created_by = ? and from_id = ? and to_id = ?", query.Get("created_by"), query.Get("from_id"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("from_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships, "created_by = ? and from_id = ?", query.Get("created_by"), query.Get("from_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships, "created_by = ? and to_id = ?", query.Get("created_by"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships, "from_id = ? and to_id = ?", query.Get("from_id"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" || query.Get("from_id") != "" || query.Get("to_id") != "" {
      ids := []string{"created_by", "from_id", "to_id"}

      var search_key string
      var search_val string

      for _, key := range ids {
        if query.Get(key) != "" {
          search_key = key
          search_val = query.Get(key)
          break
        }
      }

      // Search for query string and id
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships, search_key + " = ?", search_val).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else {
      if err := models.DB.Where(&models.Relationship{Type: searchRelationship.Type}).Find(&relationships).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": relationships})
}

// POST /relationships
// Create a relationship between two people
func CreateRelationship(c *gin.Context) {
  // Validate input
  var input CreateRelationshipInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create Relationship
  relationship := models.Relationship{CreatedBy: input.CreatedBy, FromID: input.FromID, ToID: input.ToID, Type: input.Type}
  models.DB.Create(&relationship)

  c.JSON(http.StatusOK, gin.H{"data": relationship})
}

// PATCH /relationships/:id
// Update a relationship
func UpdateRelationship(c *gin.Context) {
  // Get the relationship to be updated
  var relationship models.Relationship
  if err := models.DB.First(&relationship, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdateRelationshipInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&relationship).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": relationship})
}

// DELETE /relationships/:id
// Delete a relationship
func DeleteRelationship(c *gin.Context) {
  // Get model if exist
  var relationship models.Relationship
  if err := models.DB.First(&relationship, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&relationship)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
