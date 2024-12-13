package wallet

import (
	response "backend/src/system/responses"
	"backend/src/wallet"
	"backend/tests/wallet/fakes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeService = fakes.FakeWalletService{}
var controller = wallet.NewWalletController(&fakeService)

var responseSuccess = response.OK_WITH_DATA(&userWallet)

func Test_Controller_GetWalletByUserId_Success(t *testing.T) {
	fakeService.On("GetWalletByUserId", userId).Return(&userWallet, nil).Once()
	response := controller.GetWalletByUserId(userId)
	assert.Equal(t, responseSuccess, response)
}

func Test_Controller_GetWalletByUserId_Error(t *testing.T) {
	fakeService.On("GetWalletByUserId", userId).Return(nil, errors.New("anError")).Once()
	response := controller.GetWalletByUserId(userId)
	assert.Equal(t, 400, response.Code)
	assert.NotNil(t, response.Message)
}
