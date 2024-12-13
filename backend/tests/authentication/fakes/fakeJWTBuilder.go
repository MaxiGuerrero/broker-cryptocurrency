package fakes

import (
	"backend/src/system"

	"github.com/stretchr/testify/mock"
)

type FakeJWTBuilder struct {
	mock.Mock
}

func (f *FakeJWTBuilder) BuildToken(payload *system.Payload) string {
	args := f.Called(payload)
	return args.String(0)
}

func (f *FakeJWTBuilder) ValidateToken(tokenString string) (system.Payload, error) {
	args := f.Called(tokenString)
	if args.Get(0) == nil {
		return system.Payload{}, args.Error(1)
	}
	return args.Get(0).(system.Payload), nil
}
