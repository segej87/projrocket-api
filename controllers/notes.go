// controllers/notes.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  //"time"
  "github.com/satori/go.uuid"
  //"fmt"
  //"strings"
)

type CreateNoteInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  Text      string    `form:"text" json:"text"`
}

type UpdateNoteInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  Text      string    `form:"text" json:"text"`
}

type SearchNoteInput struct {
  Text      string    `form:"text" json:"text"`
}

// POST /notes
// Create new note
func CreateNote(c *gin.Context) {
  // Validate input
  var input CreateNoteInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create note
  note := models.Note{CreatedBy: input.CreatedBy, Text: input.Text}
  models.DB.Create(&note)

  c.JSON(http.StatusOK, gin.H{"data": note})
}

// GET /notes
// Get all notes, or find by id or query params
func FindNotes(c *gin.Context) {
  query := c.Request.URL.Query()

  var notes []models.Note

  if len(query) == 0 {
    models.DB.Find(&notes)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&notes, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else if query.Get("created_by") != "" {
    var searchNote SearchNoteInput

    if bindErr := c.BindQuery(&searchNote); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Note{Text: searchNote.Text}).Find(&notes, "created_by = ?", query.Get("created_by")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchNote SearchNoteInput

    if bindErr := c.BindQuery(&searchNote); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Note{Text: searchNote.Text}).Find(&notes).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": notes})
}

// PATCH /notes/:id
// Update a note
func UpdateNote(c *gin.Context) {
  // Get the note to be updated
  var note models.Note
  if err := models.DB.First(&note, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdateNoteInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&note).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": note})
}

// DELETE /notes/:id
// Delete a note
func DeleteNote(c *gin.Context) {
  // Get model if exist
  var note models.Note
  if err := models.DB.First(&note, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&note)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
