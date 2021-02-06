package user

type Services interface {
	CreateUser(req RegisterUserRequest) User
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(req RegisterUserRequest) User {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	newUser := s.repository.Create(user)

	return newUser
}
