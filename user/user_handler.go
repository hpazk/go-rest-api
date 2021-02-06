package user

import (
	// "net/http"

	"net/http"

	"github.com/hpazk/go-rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
}

func NewUserHandler(userService Services) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	req := new(RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(
			http.StatusBadRequest,
			"error",
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
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

	newUser := h.userService.CreateUser(*req)

	userData := UserFormatter(newUser)

	response := helper.ResponseFormatter(
		http.StatusCreated,
		"success",
		"user succesfully registered",
		userData,
	)

	return c.JSON(http.StatusCreated, response)
}
