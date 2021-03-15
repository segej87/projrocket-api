// models/meeting.go

package models

import (
  "time"
  "github.com/satori/go.uuid"
)

type Meeting struct {
  Base
  CreatedBy    uuid.UUID `form:"created_by" json:"created_by"`
  Title        string    `form:"title" json:"title"`
  Description  string    `form:"description" json:"description"`
  StartDate    time.Time `form:"start_date" json:"start_date"`
  EndDate      time.Time `form:"end_date" json:"end_date"`
  Location     string    `form:"location" json:"location"`
}
