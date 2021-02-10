package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-rest-api/app/user"
	"github.com/hpazk/go-rest-api/database"
	"github.com/hpazk/go-rest-api/routes"
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

	db := database.GetInstance()

	db.AutoMigrate(user.User{})

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.DefineApiRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
