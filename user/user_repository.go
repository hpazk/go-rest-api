package user

import (
	"github.com/jinzhu/gorm"
)

// type UsersStorage []interface{}

// var users UsersStorage

type Repository interface {
	Create(user User) (User, error)
}

// type repository struct {
// 	users *UsersStorage
// }

// func NewRepository(users *UsersStorage) *repository {
// 	return &repository{users}
// }

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {
	// users = append(users, user)

	// fmt.Println(len(users))

	// return user
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
