// models/base.go

package models

import(
  "time"
  "github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type Base struct {
  ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

