package controller

import (
	"wasabi/model"
)

// CreateUser is controller for create user
func CreateUser(c echo.Context) error {
	return c.JSON(fasthttp.StatusOK, mode.NewCreateUserResponse("0", "Success"))
}