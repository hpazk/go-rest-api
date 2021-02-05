package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Validate(user); err != nil {
			response := M{
				"code":    400,
				"status":  "error",
				"message": err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
		response := M{
			"code":    201,
			"status":  "success",
			"message": "user succesfully registered",
			"data":    user,
		}

		return c.JSON(http.StatusCreated, response)
	})

	e.GET("/users", func(c echo.Context) error {
		response := M{
			"code":    200,
			"message": "success",
		}
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
