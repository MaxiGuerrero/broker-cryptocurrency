package authentication

import (
	"backend/src/authentication/models"
	"backend/src/system"
)

type AuthRepository struct {
	database *system.Database
}

func (a AuthRepository) CreateUser(username, password, email string) *models.UserInfo {
	return nil
}
