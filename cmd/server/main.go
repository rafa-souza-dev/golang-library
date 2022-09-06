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

	// companies
	router.GET("/companies", controllers.FindAllCompanies)
	router.GET("/companies/:id", controllers.FindCompanyById)
	router.PUT("/companies/:id", controllers.UpdateCompanyById)
	router.POST("/companies", controllers.CreateCompanies)
	router.DELETE("/companies/:id", controllers.DeleteCompanyById)

	// books
	router.GET("/books", controllers.FindAllBooks)
	router.GET("/books/:id", controllers.FindBookById)
	router.PUT("/books/:id", controllers.UpdateBookById)
	router.POST("/books", controllers.CreateBooks)
	router.DELETE("/books/:id", controllers.DeleteBookById)

	router.Run("0.0.0.0:8080")
}
