package fakes

import (
	commonModels "backend/src/common/models"

	"github.com/stretchr/testify/mock"
)

type FakeWalletRepository struct {
	mock.Mock
}

func (f *FakeWalletRepository) GetWalletByUserId(userId string) *commonModels.Wallet {
	args := f.Called(userId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*commonModels.Wallet)
}

func (f *FakeWalletRepository) CreateWallet(userId string) {
	f.Called(userId)
}
