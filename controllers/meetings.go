// controllers/meetings.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "time"
  "github.com/satori/go.uuid"
)

type CreateMeetingInput struct {
  CreatedBy    uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  Title        string    `form:"title" json:"title" binding:"required"`
  Description  string    `form:"description" json:"description"`
  StartDate    time.Time `form:"start_date" json:"start_date"`
  EndDate      time.Time `form:"end_date" json:"end_date"`
  Location     string    `form:"location" json:"location"`
}

type UpdateMeetingInput struct {
  CreatedBy    uuid.UUID `form:"created_by" json:"created_by"`
  Title        string    `form:"title" json:"title"`
  Description  string    `form:"description" json:"description"`
  StartDate    time.Time `form:"start_date" json:"start_date"`
  EndDate      time.Time `form:"end_date" json:"end_date"`
  Location     string    `form:"location" json:"location"`
}

type SearchMeetingInput struct {
  Title        string    `form:"title" json:"title"`
  Description  string    `form:"description" json:"description"`
  StartDate    time.Time `form:"start_date" json:"start_date"`
  EndDate      time.Time `form:"end_date" json:"end_date"`
  Location     string    `form:"location" json:"location"`
}

// POST /meetings
// Create new meeting
func CreateMeeting(c *gin.Context) {
  // Validate input
  var input CreateMeetingInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create meeting
  meeting := models.Meeting{CreatedBy: input.CreatedBy, Title: input.Title, Description: input.Description, StartDate: input.StartDate, EndDate: input.EndDate, Location: input.Location}
  models.DB.Create(&meeting)

  c.JSON(http.StatusOK, gin.H{"data": meeting})
}

// GET /meetings
// Get all meetings, or find by id or query params
func FindMeetings(c *gin.Context) {
  query := c.Request.URL.Query()

  var meetings []models.Meeting

  if len(query) == 0 {
    models.DB.Find(&meetings)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&meetings, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else if query.Get("created_by") != "" {
    var searchMeeting SearchMeetingInput

    if bindErr := c.BindQuery(&searchMeeting); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Meeting{Title: searchMeeting.Title, StartDate: searchMeeting.StartDate, EndDate: searchMeeting.EndDate, Location: searchMeeting.Location}).Find(&meetings, "created_by = ?", query.Get("created_by")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchMeeting SearchMeetingInput

    if bindErr := c.BindQuery(&searchMeeting); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    if err := models.DB.Where(&models.Meeting{Title: searchMeeting.Title, StartDate: searchMeeting.StartDate, EndDate: searchMeeting.EndDate, Location: searchMeeting.Location}).Find(&meetings).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": meetings})
}

// PATCH /meetings/:id
// Update a meeting
func UpdateMeeting(c *gin.Context) {
  // Get the meeting to be updated
  var meeting models.Meeting
  if err := models.DB.First(&meeting, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdateMeetingInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&meeting).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": meeting})
}

// DELETE /meetings/:id
// Delete a meeting
func DeleteMeeting(c *gin.Context) {
  // Get model if exist
  var meeting models.Meeting
  if err := models.DB.First(&meeting, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&meeting)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
