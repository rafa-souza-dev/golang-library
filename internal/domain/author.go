package domain

import (
	"time"
)

type Author struct {
	Id    	  uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name	  string     `gorm:"not null" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type AuthorResponse struct {
	Name	   string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
