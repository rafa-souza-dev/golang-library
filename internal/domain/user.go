package domain

import (
	"time"
)

type User struct {
	Id    			uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  			string     `gorm:"not null" json:"name"`
	Email  			string     `gorm:"not null" json:"email"`
	Password  		string     `gorm:"not null" json:"password"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
