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

	existEmail := h.userService.CheckExistEmail(*req)

	if existEmail != nil {
		response := helper.ResponseFormatter(
			http.StatusBadRequest,
			"error",
			existEmail.Error(),
			nil,
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.CreateUser(*req)
	if err != nil {
		response := helper.ResponseFormatter(
			http.StatusInternalServerError,
			"error",
			"failed save to database",
			err.Error(),
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserFormatter(newUser)

	response := helper.ResponseFormatter(
		http.StatusCreated,
		"success",
		"user succesfully registered",
		userData,
	)

	return c.JSON(http.StatusCreated, response)
}

func (h *userHandler) GetUser(c echo.Context) error {

	response := helper.ResponseFormatter(
		http.StatusOK,
		"success",
		"login",
		nil,
	)

	return c.JSON(http.StatusCreated, response)
}
