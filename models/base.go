// models/base.go

package models

import(
  "time"
  "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
  ID        uuid.UUID    `form:"id" json:"id" gorm:"type:uuid;primaryKey"`
  CreatedAt time.Time    `form:"created_at" json:"created_at"`
  UpdatedAt time.Time    `form:"updated_at" json:"updated_at"`
}

