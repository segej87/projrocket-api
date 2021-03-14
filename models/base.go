// models/base.go

package models

import(
  "time"
  //"github.com/jinzhu/gorm"
  "github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
  ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

