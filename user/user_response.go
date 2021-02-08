package user

import "github.com/golangkit/formatime"

type UserFormat struct {
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	CreatedAt formatime.Timestamp `json:"created_at"`
	UpdatedAt formatime.Timestamp `json:"updated_at"`
}

func UserFormatter(user User) UserFormat {
	formatter := UserFormat{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return formatter
}
