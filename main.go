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

  // Relationships
  r.GET("/relationships", controllers.FindRelationships)
  r.POST("/relationships", controllers.CreateRelationship)
  r.PATCH("/relationships/:id", controllers.UpdateRelationship)
  r.DELETE("/relationships/:id", controllers.DeleteRelationship)

  // Attendances
  r.GET("/attendances", controllers.FindAttendances)
  r.POST("/attendances", controllers.CreateAttendance)
  r.PATCH("/attendances/:id", controllers.UpdateAttendance)
  r.DELETE("/attendances/:id", controllers.DeleteAttendance)

  // Documents
  r.GET("/documents", controllers.FindDocuments)
  r.POST("/documents", controllers.CreateDocument)
  r.PATCH("/documents/:id", controllers.UpdateDocument)
  r.DELETE("/documents/:id", controllers.DeleteDocument)

  r.Run()
}
