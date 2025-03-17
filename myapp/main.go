package main

import (
	"myapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	r.Run(":8080")
}
