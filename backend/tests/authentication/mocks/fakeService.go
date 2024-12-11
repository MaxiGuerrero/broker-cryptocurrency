package mocks

import (
	"backend/src/authentication/models"

	"github.com/stretchr/testify/mock"
)

type FakeService struct {
	mock.Mock
}

func (f *FakeService) Register(username, password, email string) (*models.UserInfo, error) {
	args := f.Called(username, password, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserInfo), nil
}
