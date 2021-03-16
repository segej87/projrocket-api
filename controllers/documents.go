// controllers/documents.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "github.com/satori/go.uuid"
)

type CreateDocumentInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" binding:"required"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" binding:"required"`
  Type      string    `form:"type" json:"type"`
  Origin    bool      `form:"origin" json:"origin" default:"true"`
}

type UpdateDocumentInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  FromID    uuid.UUID `form:"from_id" json:"from_id"`
  ToID      uuid.UUID `form:"to_id" json:"to_id"`
  Type      string    `form:"type" json:"type"`
  Origin    bool      `form:"origin" json:"origin"`
}

type SearchDocumentInput struct {
  Type   string `form:"type" json:"type"`
  Origin bool   `form:origin" json:"origin"`
}

// GET /documents
// Get all documents, or find by id or query params
func FindDocuments(c *gin.Context) {
  query := c.Request.URL.Query()

  var documents []models.Document

  if len(query) == 0 {
    models.DB.Find(&documents)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&documents, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchDocument SearchDocumentInput

    // Validate input
    if bindErr := c.BindQuery(&searchDocument); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    // TODO: try to fix this crazy block
    if query.Get("created_by") != "" && query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents, "created_by = ? and from_id = ? and to_id = ?", query.Get("created_by"), query.Get("from_id"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("from_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents, "created_by = ? and from_id = ?", query.Get("created_by"), query.Get("from_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents, "created_by = ? and to_id = ?", query.Get("created_by"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents, "from_id = ? and to_id = ?", query.Get("from_id"), query.Get("to_id")).Error; err != nil {
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
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents, search_key + " = ?", search_val).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else {
      if err := models.DB.Where(&models.Document{Type: searchDocument.Type}).Find(&documents).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": documents})
}

// POST /documents
// Create a document between two people
func CreateDocument(c *gin.Context) {
  // Validate input
  var input CreateDocumentInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create Document
  document := models.Document{CreatedBy: input.CreatedBy, FromID: input.FromID, ToID: input.ToID, Type: input.Type}
  models.DB.Create(&document)

  c.JSON(http.StatusOK, gin.H{"data": document})
}

// PATCH /documents/:id
// Update a document
func UpdateDocument(c *gin.Context) {
  // Get the document to be updated
  var document models.Document
  if err := models.DB.First(&document, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdateDocumentInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&document).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": document})
}

// DELETE /documents/:id
// Delete a document
func DeleteDocument(c *gin.Context) {
  // Get model if exist
  var document models.Document
  if err := models.DB.First(&document, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&document)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
