package controller

import (
	"wasabi/model"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// RegisterUser is controller for register user
func RegisterUser(c echo.Context) error {
	return c.JSON(fasthttp.StatusOK, model.NewCreateUserResponse("0", "Success"))
}

// LoginUser is controller for log-in user
func LoginUser(c echo.Context) error {
	return c.JSON(fasthttp.StatusOK, model.NewCreateUserResponse("0", "Success"))
}
