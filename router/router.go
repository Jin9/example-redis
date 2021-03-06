package router

import (
	"wasabi/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valyala/fasthttp"
)

// Init is router initial function
func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{fasthttp.MethodGet, fasthttp.MethodPost},
	}))

	v1 := e.Group("api/v1")
	{
		v1.POST("/register", controller.RegisterUser)
		v1.POST("/login", controller.LoginUser)
		v1.POST("/logout", controller.LogOutUser)
		v1.POST("/getdata", controller.ExampleData)
		// v1.GET("/file", controller.ExampleFile)
	}

	return e
}
