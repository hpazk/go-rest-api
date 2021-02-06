package user

import "fmt"

type UsersStorage []interface{}

var users UsersStorage

type Repository interface {
	Create(user User) User
}

type repository struct {
	users *UsersStorage
}

func NewRepository(users *UsersStorage) *repository {
	return &repository{users}
}

func (r *repository) Create(user User) User {
	users = append(users, user)

	fmt.Println(len(users))

	return user
}
