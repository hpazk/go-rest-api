package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/users", func(c echo.Context) error {
		response := M{
			"code":    200,
			"message": "success",
		}
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
