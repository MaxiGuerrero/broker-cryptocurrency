package models

// Struct that represent a request to create an user.
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=4,max=32"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
}
