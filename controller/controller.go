package controller

import (
	"wasabi/model"
	"wasabi/service"

	"errors"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

// CustomValidator is represent model of custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is used for validate request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validateRequest(c echo.Context, request interface{}) error {
	c.Echo().Validator = &CustomValidator{validator: validator.New()}
	if err := c.Bind(request); err != nil {
		log.Println(err.Error())
		return errors.New("Cannot binding your request")
	}
	if err := c.Validate(request); err != nil {
		log.Println(err.Error())
		return errors.New("Something wrong with your request")
	}
	return nil
}

// RegisterUser is controller for register user
func RegisterUser(c echo.Context) error {
	user := new(model.RegisterUserRequest)
	if err := validateRequest(c, user); err != nil {
		return c.JSON(fasthttp.StatusBadRequest, model.NewErrorResponse("2002", err.Error()))
	}
	if err := service.RegisterUser(user); err != nil {
		return c.JSON(fasthttp.StatusInternalServerError, model.NewErrorResponse("2002", err.Error()))
	}
	return c.JSON(fasthttp.StatusOK, model.NewCreateUserResponse("0", "Success"))
}

// LoginUser is controller for log-in user
func LoginUser(c echo.Context) error {
	user := new(model.LoginUserRequest)
	if err := validateRequest(c, user); err != nil {
		return c.JSON(fasthttp.StatusBadRequest, model.NewErrorResponse("2002", err.Error()))
	}
	atoken, err := service.LoginUser(user)
	if err != nil {
		return c.JSON(fasthttp.StatusBadRequest, model.NewErrorResponse("2002", err.Error()))
	}
	return c.JSON(fasthttp.StatusOK, model.NewLoginUserResponse("0", "Success", atoken))
}
