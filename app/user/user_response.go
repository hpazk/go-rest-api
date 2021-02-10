package user

import (
	"github.com/golangkit/formatime"
)

type UserFormat struct {
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	AuthToken string              `json:"auth_token"`
	CreatedAt formatime.Timestamp `json:"created_at"`
	UpdatedAt formatime.Timestamp `json:"updated_at"`
}

func UserFormatter(user User, auth_token string) UserFormat {
	formatter := UserFormat{
		Name:      user.Name,
		Email:     user.Email,
		AuthToken: auth_token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return formatter
}
