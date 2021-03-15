// main.go

package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/segej87/projrocket-api/models"
  "github.com/segej87/projrocket-api/controllers"
)

func main() {
  r := gin.Default()

  // Define a route to the root url simply returning a welcome message
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "welcome to projrocket api!"})
  })

  // Connect to the database
  models.ConnectDatabase()

  // Users
  r.GET("/users", controllers.FindUsers)
  r.POST("/users", controllers.CreateUser)

  // People
  r.GET("/people", controllers.FindPeople)
  r.POST("/people", controllers.CreatePerson)
  r.PATCH("/people/:id", controllers.UpdatePerson)
  r.DELETE("/people/:id", controllers.DeletePerson)

  // Meetings
  r.GET("/meetings", controllers.FindMeetings)
  r.POST("/meetings", controllers.CreateMeeting)
  r.PATCH("/meetings/:id", controllers.UpdateMeeting)
  r.DELETE("/meetings/:id", controllers.DeleteMeeting)

  // Notes
  r.GET("/notes", controllers.FindNotes)
  r.POST("/notes", controllers.CreateNote)
  r.PATCH("/notes/:id", controllers.UpdateNote)
  r.DELETE("/notes/:id", controllers.DeleteNote)

  r.Run()
}
