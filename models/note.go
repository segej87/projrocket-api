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
