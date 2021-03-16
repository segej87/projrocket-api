// models/note.go

package models

import (
  "github.com/satori/go.uuid"
)

type Note struct {
  Base
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  Text      string    `form:"text" json:"text"`
}

type Document struct {
  Base
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" gorm:"type:uuid"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" gorm:"type:uuid"`
  Type      string    `form:"type" json:"type"`
  Origin    bool      `form:"origin" json:"origin" default:"true"`
}
