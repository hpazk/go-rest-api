package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-rest-api/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	userRepository := user.NewRepository(&user.UsersStorage{})
	userService := user.NewService(userRepository)
	userHandler := user.NewUserHandler(userService)

	e.POST("/users", userHandler.RegisterUser)

	e.Logger.Fatal(e.Start(":8080"))
}
