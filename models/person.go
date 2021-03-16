// models/person.go

package models

import (
  "time"
  "github.com/satori/go.uuid"
)

type Person struct {
  Base
  CreatedBy  uuid.UUID `form:"created_by" json:"created_by"`
  FirstName  string    `form:"first_name" json:"first_name"`
  LastName   string    `form:"last_name" json:"last_name"`
  Email      string    `form:"email" json:"email"`
  Phone      string    `form:"phone" json:"phone"`
  Birthday   time.Time `form:"birthday" json:"birthday"`
  Title      string    `form:"title" json:"title"`
  Department string    `form:"department" json:"department"`
  Self       bool      `form:"self" json:"self" gorm:"default:false"`
}

type Relationship struct {
  Base
  CreatedBy uuid.UUID `form:"created_by" json:"created_by"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" gorm:"type:uuid"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" gorm:"type:uuid"`
  Type      string    `form:"type" json:"type"`
}

type Attendance struct {
  Base
  CreatedBy uuid.UUID `form:"created_by" json:"created_by" gorm:"type:uuid"`
  FromID    uuid.UUID `form:"from_id" json:"from_id" gorm:"type:uuid"`
  ToID      uuid.UUID `form:"to_id" json:"to_id" gorm:"type:uuid"`
  Type      string    `form:"type" json:"type"`
}
