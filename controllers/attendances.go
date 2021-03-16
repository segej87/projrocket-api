// controllers/attendances.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "github.com/satori/go.uuid"
)

type CreateAttendanceInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by" binding:"required"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" binding:"required"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" binding:"required"`
  Type      string    `form:"type" json:"type"`
}

type UpdateAttendanceInput struct {
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  FromID    uuid.UUID `form:"from_id" json:"from_id"`
  ToID      uuid.UUID `form:"to_id" json:"to_id"`
  Type      string    `form:"type" json:"type"`
}

type SearchAttendanceInput struct {
  Type string `form:"type" json:"type"`
}

// GET /attendances
// Get all attendances, or find by id or query params
func FindAttendances(c *gin.Context) {
  query := c.Request.URL.Query()

  var attendances []models.Attendance

  if len(query) == 0 {
    models.DB.Find(&attendances)
  } else if query.Get("id") != "" {
    if err := models.DB.Find(&attendances, "id = ?", query.Get("id")).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  } else {
    var searchAttendance SearchAttendanceInput

    // Validate input
    if bindErr := c.BindQuery(&searchAttendance); bindErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
      return
    }

    // TODO: try to fix this crazy block
    if query.Get("created_by") != "" && query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances, "created_by = ? and from_id = ? and to_id = ?", query.Get("created_by"), query.Get("from_id"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("from_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances, "created_by = ? and from_id = ?", query.Get("created_by"), query.Get("from_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("created_by") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances, "created_by = ? and to_id = ?", query.Get("created_by"), query.Get("to_id")).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else if query.Get("from_id") != "" && query.Get("to_id") != "" {
      // Search for query string and ids
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances, "from_id = ? and to_id = ?", query.Get("from_id"), query.Get("to_id")).Error; err != nil {
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
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances, search_key + " = ?", search_val).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
        return
      }
    } else {
      if err := models.DB.Where(&models.Attendance{Type: searchAttendance.Type}).Find(&attendances).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }
    }
  }

  c.JSON(http.StatusOK, gin.H{"data": attendances})
}

// POST /attendances
// Create an attendance between a person and a meeting
func CreateAttendance(c *gin.Context) {
  // Validate input
  var input CreateAttendanceInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create attendance
  attendance := models.Attendance{CreatedBy: input.CreatedBy, FromID: input.FromID, ToID: input.ToID, Type: input.Type}
  models.DB.Create(&attendance)

  c.JSON(http.StatusOK, gin.H{"data": attendance})
}

// PATCH /attendances/:id
// Update an attendance
func UpdateAttendance(c *gin.Context) {
  // Get the attendance to be updated
  var attendance models.Attendance
  if err := models.DB.First(&attendance, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Validate input
  var input UpdateAttendanceInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&attendance).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": attendance})
}

// DELETE /attendances/:id
// Delete an attendance
func DeleteAttendance(c *gin.Context) {
  // Get model if exist
  var attendance models.Attendance
  if err := models.DB.First(&attendance, "id = ?", c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&attendance)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
