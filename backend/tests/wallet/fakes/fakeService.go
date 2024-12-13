package fakes

import (
	"backend/src/common/models"

	"github.com/stretchr/testify/mock"
)

type FakeWalletService struct {
	mock.Mock
}

func (f *FakeWalletService) GetWalletByUserId(userId string) (*models.Wallet, error) {
	args := f.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Wallet), nil
}
