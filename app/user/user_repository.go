package user

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	InsertUser(user User) (User, error)
	FindByEmail(email string) *User
	FindUserByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) *User {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
