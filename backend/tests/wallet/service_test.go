package wallet

import (
	"backend/src/wallet"
	"backend/tests/wallet/fakes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var walletRepository = fakes.FakeWalletRepository{}
var service = wallet.NewWalletService(&walletRepository)

func Test_Service_GetWalletByUserId_Success(t *testing.T) {
	walletRepository.On("GetWalletByUserId", userId).Return(&userWallet, nil).Once()
	wallet, err := service.GetWalletByUserId(userId)
	assert.NoError(t, err)
	assert.Equal(t, userWallet, *wallet)
}

func Test_Service_GetWalletByUserId_Error(t *testing.T) {
	walletRepository.On("GetWalletByUserId", userId).Return(nil, errors.New("error")).Once()
	wallet, err := service.GetWalletByUserId(userId)
	assert.Error(t, err)
	assert.Nil(t, wallet)
}
