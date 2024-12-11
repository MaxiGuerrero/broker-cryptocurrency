package mocks

import (
	"backend/src/authentication/models"

	"github.com/stretchr/testify/mock"
)

type FakeAuthRepository struct {
	mock.Mock
}

func (a *FakeAuthRepository) CreateUser(username, password, email string) *models.UserInfo {
	args := a.Called(username, password, email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.UserInfo)
}

func (a *FakeAuthRepository) FindUserByUsernameAndEmail(username, email string) *models.User {
	args := a.Called(username, email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.User)
}
