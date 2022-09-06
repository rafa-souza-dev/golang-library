package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"	
	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
	"time"
)

func CreateBooks(c *gin.Context) {
    var newBook domain.Book
    var author domain.Author
    var company domain.Company

    if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(400, gin.H{"message": err})
		fmt.Println(time.Now())
		fmt.Println(err)
        return
    }

	checkAuthor := database.DB_.Find(&author, newBook.AuthorId)
	if checkAuthor.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "author don't exists"})
		return
	}

	checkCompany := database.DB_.Find(&company, newBook.CompanyId)
	if checkCompany.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "company don't exists"})
		return
	}

	if newBook.Title == "" {
		c.IndentedJSON(400, gin.H{"message": "title is required"})
		return
	}
	
	book := domain.Book{
		Id: newBook.Id,
		Title: newBook.Title,
		Subtitle: newBook.Subtitle,
		AuthorId: newBook.AuthorId,
		Author: author,
		CompanyId: newBook.CompanyId,
		Company: company,
		PublicationDate: newBook.PublicationDate,
		CreatedAt: newBook.CreatedAt,
		UpdatedAt: newBook.UpdatedAt,
	}

    result := database.DB_.Create(&book)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "insert database error"})
		return
	}

	bookResponse := domain.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Subtitle: book.Subtitle,
		Author: author,
		Company: company,
		PublicationDate: book.PublicationDate,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	c.IndentedJSON(201, bookResponse)
}

func FindAllBooks(c *gin.Context) {
	var books []domain.Book
	var author domain.Author
	var company domain.Company

	result := database.DB_.Find(&books)
	if result.Error == nil {
		for i, book := range books {
			database.DB_.Find(&author, book.AuthorId)
			database.DB_.Find(&company, book.CompanyId)
			book.Author = author
			fmt.Println(book.Author)
			book.Company = company
			fmt.Println(book.Company)
			books[i] = book
		}
		c.IndentedJSON(200, books)
		return
	}
}

func FindBookById(c *gin.Context) {
	id := c.Param("id")

	var book domain.Book
	var author domain.Author
	var company domain.Company

	result := database.DB_.First(&book, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "book not found"})
		return
	}

	database.DB_.Find(&author, book.AuthorId)
	database.DB_.Find(&company, book.CompanyId)

	bookResponse := domain.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Subtitle: book.Subtitle,
		Author: author,
		Company: company,
		PublicationDate: book.PublicationDate,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	c.IndentedJSON(200, bookResponse)
}

func UpdateBookById(c *gin.Context) {
	id := c.Param("id")

	var book domain.Book
	var author domain.Author
	var company domain.Company

	result := database.DB_.First(&book, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "book not found"})
		return
	}

	var updatedBook domain.Book

    if err := c.BindJSON(&updatedBook); err != nil {
        return
    }

	book.Title = updatedBook.Title
	book.Subtitle = updatedBook.Subtitle
	book.AuthorId = updatedBook.AuthorId
	book.CompanyId = updatedBook.CompanyId
	book.PublicationDate = updatedBook.PublicationDate

	save := database.DB_.Save(&book)
	if save.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "failed to edit book"})
		return
	}

	database.DB_.Find(&author, book.AuthorId)
	database.DB_.Find(&company, book.CompanyId)

	book.Author = author
	book.Company = company

	c.IndentedJSON(200, book)
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")

	var book domain.Book

	result := database.DB_.First(&book, id)
	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "book not found"})
		return
	}

	delete := database.DB_.Delete(&book)
	if delete.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "could not delete this book"})
		return
	}

	c.IndentedJSON(204, gin.H{"message": "book successfully deleted"})
}
