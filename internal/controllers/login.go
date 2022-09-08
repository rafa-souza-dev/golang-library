package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
	"github.com/rafa-souza-dev/library/internal/services"
)

func Login(c *gin.Context) {
    var loginData domain.Login

    if err := c.BindJSON(&loginData); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid login data"})
        return
    }

	if loginData.Email == "" || loginData.Password == "" {
		c.IndentedJSON(400, gin.H{"message": "inform all data"})
		return
	}

	var user domain.User

	hasUserError := database.DB_.Where("email = ?", loginData.Email).First(&user).Error
	if hasUserError != nil {
		c.IndentedJSON(400, gin.H{"message": "cannot find user"})
		return
	}

	if user.Password != services.SHA256Encoder(loginData.Password) {
		c.IndentedJSON(400, gin.H{"message": "invalid user credentials"})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.Id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
