package user

import (
	"github.com/hpazk/go-rest-api/auth"
	"github.com/hpazk/go-rest-api/database"
	"github.com/hpazk/go-rest-api/helper"
	"github.com/hpazk/go-rest-api/middleware"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Routes() []helper.Route {
	db := database.GetInstance()

	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	authService := auth.NewService()

	userHandler := NewUserHandler(userService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.RegisterUser,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: userHandler.LoginUser,
		},
		{
			Method:     echo.GET,
			Path:       "/users",
			Handler:    userHandler.GetUser,
			Middleware: []echo.MiddlewareFunc{middleware.JwtMiddleWare()},
		},
	}
}
