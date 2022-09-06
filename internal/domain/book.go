package domain

import (
	"time"
)

type Book struct {
	Id    	         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title            string     `gorm:"not null" json:"title"`
	Subtitle         string     `json:"subtitle"`
	AuthorId         uint       `gorm:"not null" json:"author_id"`
	Author           Author     `json:"author"`
	CompanyId        uint       `gorm:"not null" json:"company_id"`
	Company          Company    `json:"company"`
	PublicationDate  time.Time  `json:"publication_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type BookResponse struct {
	Id    	         uint       	  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title            string           `gorm:"not null" json:"title"`
	Subtitle         string           `json:"subtitle"`
	Author           Author           `json:"author"`
	Company          Company          `json:"company"`
	PublicationDate  time.Time  	  `json:"publication_date"`
	CreatedAt        time.Time  	  `json:"created_at"`
	UpdatedAt        time.Time  	  `json:"updated_at"`
}
