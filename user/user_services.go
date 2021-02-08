package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	CreateUser(req RegisterUserRequest) (User, error)
	CheckExistEmail(req RegisterUserRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(req RegisterUserRequest) (User, error) {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHashed)

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) CheckExistEmail(req RegisterUserRequest) error {
	email := req.Email

	if user := s.repository.FindEmail(email); user != nil {
		return errors.New("email already regietered")
	}

	return nil
}
