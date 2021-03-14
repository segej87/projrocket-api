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
  
  models.ConnectDatabase()
  
  r.GET("/people", controllers.FindBooks)

  r.Run()
}

