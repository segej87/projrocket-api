// models/user.go

package models

type User struct {
  Base
  Username  string    `form:"username" json:"username"`
}

