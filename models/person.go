// models/person.go

package models

import (
  "time"
  "github.com/satori/go.uuid"
)

type Person struct {
  Base
  CreatedBy  uuid.UUID `form:"created_by" json:"created_by"`
  FirstName  string    `form:"firstname" json:"firstname"`
  LastName   string    `form:"lastname" json:"lastname"`
  Email      string    `form:"email" json:"email"`
  Phone      string    `form:"phone" json:"phone"`
  Birthday   time.Time `form:"birthday" json:"birthday"`
  Title      string    `form:"title" json:"title"`
  Department string    `form:"department" json:"department"`
  Self       bool      `form:"self" json:"self" gorm:"default:false"`
}
