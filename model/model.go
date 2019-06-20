package model

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
