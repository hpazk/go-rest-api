package routes

import (
	"github.com/hpazk/go-rest-api/helper"
	"github.com/hpazk/go-rest-api/user"
	"github.com/labstack/echo/v4"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handlers{
		user.UserRoutes{},
	}

	var routes []helper.Route
	for _, controller := range handlers {
		routes = append(routes, controller.Routes()...)
	}

	api := e.Group("/api/v1")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}
