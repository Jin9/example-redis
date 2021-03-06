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

// RegisterUserResponse is represent model of CreateUser Response
type RegisterUserResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}

// NewRegisterUserResponse used for create `RegisterUserResponse`
func NewRegisterUserResponse(statusCode string, statusDesc string) *RegisterUserResponse {
	return &RegisterUserResponse{
		StatusCode: statusCode,
		StatusDesc: statusDesc,
	}
}

// User is represent model of User
type User struct {
	UToken string `json:"uToken"`
	PToken string `json:"pToken"`
}

// NewUser is for allocate new user model
func NewUser(uToken string, ptoken string) *User {
	return &User{
		UToken: uToken,
		PToken: ptoken,
	}
}

// UserToken is represent of model UToken in redis db
type UserToken struct {
	Username string `json:"username"`
	AToken   string `json:"atoken"`
}

// NewUserToken is used for create new model of UserToken
func NewUserToken(username string, atoken string) *UserToken {
	return &UserToken{
		Username: username,
		AToken:   atoken,
	}
}

// AccessToken is represent model of Access token
type AccessToken struct {
	AToken     string
	HashAToken string
}

// NewAccessToken is used for create AccessToken model
func NewAccessToken(atoken string, hashAtoken string) *AccessToken {
	return &AccessToken{
		AToken:     atoken,
		HashAToken: hashAtoken,
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

// ExampleRequest is
type ExampleRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewExampleRequest is
func NewExampleRequest(username string, password string) *ExampleRequest {
	return &ExampleRequest{
		Username: username,
		Password: password,
	}
}
