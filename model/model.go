package model

// ErrorResponse is represent model of error response
type ErrorResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}

// NewErrorResponse is used for allocate new Error response model
func NewErrorResponse(statusCode string, statusDesc string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		StatusDesc: statusDesc,
	}
}

// RegisterUserRequest is represent model of CreateUser Request
type RegisterUserRequest struct {
	UserName  string `json:"username" validate:"required"`
	PToken    string `json:"ptoken" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

// CreateUserResponse is represent model of CreateUser Response
type CreateUserResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
	UToken     string `json:"utoken"`
}

// NewCreateUserResponse used for create `CreateUserResponse`
func NewCreateUserResponse(statusCode string, statusDesc string) *CreateUserResponse {
	return &CreateUserResponse{
		StatusCode: statusCode,
		StatusDesc: statusDesc,
	}
}

// User is represent model of User
type User struct {
	UToken    string `json:"uToken"`
	PToken    string `json:"pToken"`
	IsLogedIn bool   `json:"isLogedIn"`
}

// NewUser is for allocate new user model
func NewUser(uToken string, ptoken string, isLogedIn bool) *User {
	return &User{
		UToken:    uToken,
		PToken:    ptoken,
		IsLogedIn: isLogedIn,
	}
}

// LoginUserRequest is represent Login Request
type LoginUserRequest struct {
	Username string `json:"username"`
	PToken   string `json:"ptoken"`
}

// LoginUserResponse is represent Login Response
type LoginUserResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
	AToken     string `json:"aToken"`
}

// NewLoginUserResponse is used for create new Login User Response
func NewLoginUserResponse(statusCode string, statusDesc string, aToken string) *LoginUserResponse {
	return &LoginUserResponse{
		StatusCode: statusCode,
		StatusDesc: statusDesc,
		AToken:     aToken,
	}
}
