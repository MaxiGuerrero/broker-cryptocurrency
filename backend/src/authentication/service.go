package authentication

import "backend/src/authentication/interfaces"

type AuthService struct {
	AuthRepository interfaces.IAuthRepository
}

func (a AuthService) Register(username, password, email string) error {
	return nil
}
