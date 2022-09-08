package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
	"github.com/rafa-souza-dev/library/internal/services"
)

func CreateUser(c *gin.Context) {
    var newUser domain.User

    if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid user data"})
        return
    }

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		c.IndentedJSON(400, gin.H{"message": "inform all data"})
		return
	}

	newUser.Password = services.SHA256Encoder(newUser.Password)

    result := database.DB_.Create(&newUser)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "insert in database error"})
		return
	}

	c.Status(201)
}
