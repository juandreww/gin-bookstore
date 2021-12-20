package main

import (
  "github.com/gin-gonic/gin"
  "github.com/rahmanfadhil/gin-bookstore/models"
  "github.com/rahmanfadhil/gin-bookstore/controllers"
  "C"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()
  r.GET("/books", controllers.FindBooks)
  r.POST("/books", controllers.CreateBook)

  r.Run()
}