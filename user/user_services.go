package user

type Services interface {
	CreateUser(req RegisterUserRequest) (User, error)
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
	user.Password = req.Password

	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
