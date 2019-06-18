package controller

import (
	"wasabi/model"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// CreateUser is controller for create user
func CreateUser(c echo.Context) error {
	return c.JSON(fasthttp.StatusOK, model.NewCreateUserResponse("0", "Success"))
}
