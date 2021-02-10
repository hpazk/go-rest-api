package user

import (
	// "net/http"

	"net/http"

	"github.com/hpazk/go-rest-api/auth"
	"github.com/hpazk/go-rest-api/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService Services
	authService auth.Service
}

func NewUserHandler(userService Services, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
			errMessage,
			nil,
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
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	auth_token, err := h.authService.GetAccessToken(newUser.ID)
	if err != nil {
		response := helper.ResponseFormatter(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}
	userData := UserFormatter(newUser, auth_token)

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
		"get user",
		nil,
	)

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c echo.Context) error {
	req := new(LoginUserRequest)
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
			errMessage,
			nil,
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	userAuth, err := h.userService.AuthUser(*req)
	if err != nil {
		// errors := helper.ErrorFormatter(err)
		// errMessage := helper.M{"errors": errors}

		response := helper.ResponseFormatter(
			http.StatusUnauthorized,
			"error",
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusUnauthorized, response)
	}

	auth_token, err := h.authService.GetAccessToken(userAuth.ID)
	if err != nil {
		response := helper.ResponseFormatter(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := UserFormatter(userAuth, auth_token)

	response := helper.ResponseFormatter(
		http.StatusOK,
		"success",
		"user authenticated",
		userData,
	)

	// return c.JSON(http.StatusCreated, response)

	return c.JSON(http.StatusOK, response)

}
