package authentication

import (
	"backend/src/authentication/interfaces"
	"backend/src/authentication/models"
	"errors"
)

type AuthService struct {
	encrypter      interfaces.IEncrypter
	authRepository interfaces.IAuthRepository
	jwtBuilder     interfaces.IJWTBuilder
}

func (a *AuthService) Register(username, password, email string) (*models.UserInfo, error) {
	userFound := a.authRepository.FindUserByUsernameAndEmail(username, email)
	if userFound == nil {
		return nil, errors.New("username and/or email already exists")
	}
	passwordHashed := a.encrypter.GenerateHash(password)
	userInfo := a.authRepository.CreateUser(username, passwordHashed, email)
	return userInfo, nil
}
