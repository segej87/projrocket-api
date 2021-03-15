// models/setup.go

package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/satori/go.uuid"
)

var DB *gorm.DB

func (base *Base) BeforeCreate(scope *gorm.Scope) (err error) {
  base.ID = uuid.NewV4()

  return
}

func ConnectDatabase() {
  database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&User{})
  database.AutoMigrate(&Person{})
  database.AutoMigrate(&Meeting{})
  database.AutoMigrate(&Note{})

  DB = database
}
