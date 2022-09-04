package domain

import (
	"time"
)

type Company struct {
	Id    			uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  			string     `gorm:"not null" json:"name"`
	FoundationDate  time.Time  `json:"foundation_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type CompanyResponse struct {
	Name	        string     `json:"name"`
	FoundationDate  time.Time  `json:"foundation_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
