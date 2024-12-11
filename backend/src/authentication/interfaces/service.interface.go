package interfaces

import "backend/src/authentication/models"

type IAuthService interface {
	Register(username, password, email string) (*models.UserInfo, error)
	Login(username, password string) (*string, error)
}
