package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
)

func CreateAuthors(c *gin.Context) {
    var newAuthor domain.Author

    if err := c.BindJSON(&newAuthor); err != nil {
        return
    }

	if newAuthor.Name == "" {
		c.IndentedJSON(404, gin.H{"message": "name is required"})
		return
	}
	
	author := domain.Author{
		Name: newAuthor.Name,
		CreatedAt: newAuthor.CreatedAt,
		UpdatedAt: newAuthor.UpdatedAt,
	}

    result := database.DB_.Create(&author)
	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "invalid author data"})
		return
	}

	c.IndentedJSON(201, newAuthor)
}

func FindAllAuthors(c *gin.Context) {
	var authors []domain.Author
	result := database.DB_.Find(&authors)
	if result.Error == nil {
		c.IndentedJSON(200, authors)
		return
	}
}

func FindAuthorById(c *gin.Context) {
	id := c.Param("id")

	var author domain.Author

	result := database.DB_.First(&author, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "author not found"})
		return
	}

	c.IndentedJSON(200, author)
}

func UpdateAuthorById(c *gin.Context) {
	id := c.Param("id")

	var author domain.Author

	result := database.DB_.First(&author, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "author not found"})
		return
	}

	var updatedAuthor domain.Author

    if err := c.BindJSON(&updatedAuthor); err != nil {
        return
    }

	if updatedAuthor.Name == "" {
		c.IndentedJSON(404, gin.H{"message": "name is required"})
		return
	}

	author.Name = updatedAuthor.Name

	save := database.DB_.Save(&author)
	if save.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "failed to edit author"})
		return
	}

	c.IndentedJSON(200, author)
}

func DeleteAuthorById(c *gin.Context) {
	id := c.Param("id")

	var author domain.Author

	result := database.DB_.First(&author, id)
	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "author not found"})
		return
	}

	delete := database.DB_.Delete(&author)
	if delete.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "could not delete this author"})
		return
	}

	c.IndentedJSON(204, gin.H{"message": "author successfully deleted"})
}
