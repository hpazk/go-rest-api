package user

type UserFormat struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserFormatter(user User) UserFormat {
	formatter := UserFormat{
		Name:  user.Name,
		Email: user.Email,
	}

	return formatter
}
