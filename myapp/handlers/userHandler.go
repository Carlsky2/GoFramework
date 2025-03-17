package handlers

import (
	"myapp/models"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users = append(users, user)
	c.JSON(201, user)
}
