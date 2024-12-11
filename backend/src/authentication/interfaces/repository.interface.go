package interfaces

import "backend/src/authentication/models"

type IAuthRepository interface {
	CreateUser(username, password, email string) *models.UserInfo
	FindUserByUsername(username string) *models.User
}
