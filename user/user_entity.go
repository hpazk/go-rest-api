package user

import "gorm.io/gorm"

// struct yang berhubungan dengan database
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
