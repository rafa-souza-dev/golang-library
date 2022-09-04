package main

import (
	"github.com/rafa-souza-dev/library/internal/domain"
	"github.com/rafa-souza-dev/library/internal/database"
)

func main() {
	database.ConnectionDB()

	database.DB_.Migrator().CreateTable(&domain.Author{})
	database.DB_.Migrator().CreateTable(&domain.Company{})
	database.DB_.Migrator().CreateTable(&domain.Book{})
}
