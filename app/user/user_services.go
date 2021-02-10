package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	CreateUser(req RegisterUserRequest) (User, error)
	CheckExistEmail(req RegisterUserRequest) error
	AuthUser(req LoginUserRequest) (User, error)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) CheckExistEmail(req RegisterUserRequest) error {
	email := req.Email

	if existUser := s.repository.FindByEmail(email); existUser != nil {
		return errors.New("email already registered")
	}

	return nil
}

func (s *service) AuthUser(req LoginUserRequest) (User, error) {
	email := req.Email
	password := req.Password

	// var user User

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, errors.New("email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}
