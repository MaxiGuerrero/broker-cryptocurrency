package fakes

import (
	"github.com/stretchr/testify/mock"
)

type FakeEncrypter struct {
	mock.Mock
}

func (e *FakeEncrypter) GenerateHash(password string) string {
	args := e.Called(password)
	return args.String(0)
}

func (e *FakeEncrypter) Compare(hashedPassword, password string) bool {
	args := e.Called(hashedPassword, password)
	return args.Bool(0)
}
