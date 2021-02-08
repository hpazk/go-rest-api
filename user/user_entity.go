package user

import "github.com/hpazk/go-rest-api/helper"

// struct yang berhubungan dengan database
type User struct {
	helper.BaseModel
	Name     string
	Email    string
	Password string
}
