package authentication

import (
	"backend/src/authentication/interfaces"
	"backend/src/authentication/models"
)

type AuthService struct {
	authRepository interfaces.IAuthRepository
}

func (a AuthService) Register(username, password, email string) *models.UserInfo {
	return nil
}
