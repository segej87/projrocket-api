// models/person.go

package models

type Person struct {
  Base
  FirstName string `form: "firstname" json:"firstname"`
  LastName  string `form: "lastname" json:"lastname"`
  Title     string `form: "title" json:"title"`
}

