package main

import (
	"github.com/go-playground/validator/v10"
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

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.DefineApiRoutes(e)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		if req.URL.Path[:4] == "/api" {
			e.ServeHTTP(res, req)
		}

		return
	})

	server.Logger.Fatal(server.Start(":8080"))
}
