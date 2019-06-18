package model

// CreateUserResponse is represent model of a Response for api CreateUser
type CreateUserResponse struct{
	Status string `json:"status"`
	StatusDesc string `json:"statusDesc"`
}