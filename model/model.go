package model

// User is
type User struct {
}

// CreateUserResponse is represent model of a Response for api CreateUser
type CreateUserResponse struct {
	Status     string `json:"status"`
	StatusDesc string `json:"statusDesc"`
}

// NewCreateUserResponse used for create `CreateUserResponse`
func NewCreateUserResponse(status string, statusDesc string) *CreateUserResponse {
	return &CreateUserResponse{
		Status:     status,
		StatusDesc: statusDesc,
	}
}
