package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		response := M{
			"code":    200,
			"message": "success",
		}
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
