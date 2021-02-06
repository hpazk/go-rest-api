package user

import (
	"github.com/hpazk/go-rest-api/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (controller UserRoutes) Routes() []helper.Route {
	userRepository := NewRepository(&UsersStorage{})
	userService := NewService(userRepository)
	userHandler := NewUserHandler(userService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.RegisterUser,
		},
	}
}
