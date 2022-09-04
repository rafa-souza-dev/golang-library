package controllers

import (
	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
	"github.com/gin-gonic/gin"
)

func CreateCompanies(c *gin.Context) {
    var newCompany domain.Company

    if err := c.BindJSON(&newCompany); err != nil {
        return
    }

	if newCompany.Name == "" {
		c.IndentedJSON(404, gin.H{"message": "name is required"})
		return
	}
	
	company := domain.Company{
		Name: newCompany.Name,
		FoundationDate: newCompany.FoundationDate,
		CreatedAt: newCompany.CreatedAt,
		UpdatedAt: newCompany.UpdatedAt,
	}

    result := database.DB_.Create(&company)
	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "invalid company data"})
		return
	}

	c.IndentedJSON(201, newCompany)
}

func FindAllCompanies(c *gin.Context) {
	var companies []domain.Company
	result := database.DB_.Find(&companies)
	if result.Error == nil {
		c.IndentedJSON(200, companies)
		return
	}
}

func FindCompanyById(c *gin.Context) {
	id := c.Param("id")

	var company domain.Company

	result := database.DB_.First(&company, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "company not found"})
		return
	}

	c.IndentedJSON(200, company)
}

func UpdateCompanyById(c *gin.Context) {
	id := c.Param("id")

	var company domain.Company

	result := database.DB_.First(&company, id)

	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "company not found"})
		return
	}

	var updatedCompany domain.Company

    if err := c.BindJSON(&updatedCompany); err != nil {
        return
    }

	if updatedCompany.Name == "" {
		c.IndentedJSON(404, gin.H{"message": "name is required"})
		return
	}

	company.Name = updatedCompany.Name
	company.FoundationDate = updatedCompany.FoundationDate

	save := database.DB_.Save(&company)
	if save.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "failed to edit company"})
		return
	}

	c.IndentedJSON(200, company)
}

func DeleteCompanyById(c *gin.Context) {
	id := c.Param("id")

	var company domain.Company

	result := database.DB_.First(&company, id)
	if result.Error != nil {
		c.IndentedJSON(404, gin.H{"message": "company not found"})
		return
	}

	delete := database.DB_.Delete(&company)
	if delete.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "could not delete this company"})
		return
	}

	c.IndentedJSON(204, gin.H{"message": "company successfully deleted"})
}
