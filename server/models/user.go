package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"not null;unique"`
	Email  string `json:"email" gorm:"size:255;not null;unique"`
	Phone  int `json:"phone" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null;"`
	Role  string `json:"role" gorm:"gorm:"type:enum('client', 'admin','content manager')" not null;DEFAULT:client"`
	verified  bool `json:"verified"`
	CreatedAt time.Time
	UpdatedAt time.Time
}