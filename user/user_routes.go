package user

import (
	"github.com/hpazk/go-rest-api/database"
	"github.com/hpazk/go-rest-api/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Routes() []helper.Route {
	db := database.GetInstance()

	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	userHandler := NewUserHandler(userService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.RegisterUser,
		},
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetUser,
		},
	}
}
