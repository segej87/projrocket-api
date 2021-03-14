// models/person.go

package models

type Person struct {
  Base
  FirstName string `json:"firstname"`
  LastName  string `json:"lastname"`
  Title     string `json:"title"`
}

