package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hpazk/go-rest-api/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

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

	e.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			response := helper.ResponseFormatter(
				http.StatusBadRequest,
				"error",
				err.Error(),
				nil,
			)
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := c.Validate(user); err != nil {
			errors := helper.ErrorFormatter(err)
			errMessage := helper.M{"errors": errors}

			response := helper.ResponseFormatter(
				http.StatusBadRequest,
				"error",
				"registering user failed",
				errMessage,
			)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := helper.ResponseFormatter(
			http.StatusCreated,
			"success",
			"user succesfully registered",
			user,
		)

		return c.JSON(http.StatusCreated, response)
	})

	e.GET("/users", func(c echo.Context) error {
		response := helper.ResponseFormatter(
			http.StatusCreated,
			"success",
			"user succesfully registered",
			nil,
		)
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
