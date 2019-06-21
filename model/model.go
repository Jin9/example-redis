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
	Status     string `json:"status"`
	StatusDesc string `json:"statusDesc"`
	UToken     string `json:"utoken"`
}

// NewCreateUserResponse used for create `CreateUserResponse`
func NewCreateUserResponse(status string, statusDesc string) *CreateUserResponse {
	return &CreateUserResponse{
		Status:     status,
		StatusDesc: statusDesc,
	}
}

// User is represent model of User
type User struct {
	UToken         string `json:"uToken"`
	AccessExpireIn int64  `json:"accessExpireIn"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

// NewUser is for allocate new user model
func NewUser(uToken string, accessExpireIn int64, email string, phone string) *User {
	return &User{
		UToken:         uToken,
		AccessExpireIn: accessExpireIn,
		Email:          email,
		Phone:          phone,
	}
}
