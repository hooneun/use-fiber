package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
}
