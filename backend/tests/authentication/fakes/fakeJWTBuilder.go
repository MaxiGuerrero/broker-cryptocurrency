package fakes

import (
	"backend/src/authentication/models"

	"github.com/stretchr/testify/mock"
)

type FakeJWTBuilder struct {
	mock.Mock
}

func (f *FakeJWTBuilder) BuildToken(payload *models.Payload) string {
	args := f.Called(payload)
	return args.String(0)
}

func (f *FakeJWTBuilder) ValidateToken(tokenString string) (*models.Payload, error) {
	args := f.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Payload), nil
}
