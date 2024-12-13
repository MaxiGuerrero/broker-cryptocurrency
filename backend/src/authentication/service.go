package authentication

import (
	"backend/src/authentication/interfaces"
	"backend/src/authentication/models"
	"backend/src/system"
	"errors"
)

type AuthService struct {
	encrypter      interfaces.IEncrypter
	authRepository interfaces.IAuthRepository
	jwtBuilder     interfaces.IJWTBuilder
}

func NewAuthService(encrypter interfaces.IEncrypter, authRepository interfaces.IAuthRepository, jwtBuilder interfaces.IJWTBuilder) *AuthService {
	return &AuthService{encrypter, authRepository, jwtBuilder}
}

func (a *AuthService) Register(username, password, email string) (*models.UserInfo, error) {
	userFound := a.authRepository.FindUserByUsername(username)
	if userFound != nil {
		return nil, errors.New("username and/or email already exists")
	}
	passwordHashed := a.encrypter.GenerateHash(password)
	userInfo := a.authRepository.CreateUser(username, passwordHashed, email)
	return userInfo, nil
}

func (a *AuthService) Login(username, password string) (*string, error) {
	userFound := a.authRepository.FindUserByUsername(username)
	if userFound == nil {
		return nil, errors.New("username and/or password are incorrect")
	}
	if userFound.Status == models.BLOCKED || userFound.Status == models.INACTIVE {
		return nil, errors.New("username and/or password are incorrect")
	}
	if ok := a.encrypter.Compare(userFound.Password, password); !ok {
		return nil, errors.New("username and/or password are incorrect")
	}
	token := a.jwtBuilder.BuildToken(&system.Payload{
		UserId:   userFound.ID,
		Username: userFound.Username,
		Role:     userFound.Role.String(),
	})
	return &token, nil
}
