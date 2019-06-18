package controller

// CreateUser is controller for create user
func CreateUser(c echo.Context) error {
	return c.JSON(fasthttp.StatusOK, "Hello, World")
}