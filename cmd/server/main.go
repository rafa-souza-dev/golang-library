package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rafa-souza-dev/library/internal/controllers"
	"github.com/rafa-souza-dev/library/internal/database"
)

func main() {
	// database conn
	database.ConnectionDB()

	router := gin.Default()

	// authors
	router.GET("/authors", controllers.FindAllAuthors)
	router.GET("/authors/:id", controllers.FindAuthorById)
	router.PUT("/authors/:id", controllers.UpdateAuthorById)
	router.POST("/authors", controllers.CreateAuthors)
	router.DELETE("/authors/:id", controllers.DeleteAuthorById)

	router.Run("0.0.0.0:8080")
}
