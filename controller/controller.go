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
	return c.JSON(fasthttp.StatusOK, model.NewRegisterUserResponse("0", "Success"))
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

// LogOutUser is controller for log-out user
func LogOutUser(c echo.Context) error {
	atoken := c.Request().Header.Get("accessToken")
	if err := service.LogOutUser(atoken); err != nil {
		return c.JSON(fasthttp.StatusBadRequest, model.NewErrorResponse("2002", err.Error()))
	}
	return c.JSON(fasthttp.StatusOK, model.NewErrorResponse("0", "Success"))
}

// ExampleData is used for send example get request
func ExampleData(c echo.Context) error {
	exRequest := new(model.ExampleRequest)
	if err := validateRequest(c, exRequest); err != nil {
		return c.JSON(fasthttp.StatusBadRequest, model.NewErrorResponse("2002", err.Error()))
	}
	return c.JSON(fasthttp.StatusOK, model.NewExampleRequest(exRequest.Username, exRequest.Password))
}

// ExampleFile is used for send file
// func ExampleFile(c echo.Context) error {
// 	return c.File("example.txt")
// }
