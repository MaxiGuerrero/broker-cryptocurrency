package wallet

import (
	response "backend/src/system/responses"
	"backend/src/wallet/interfaces"
)

type WalletController struct {
	IWalletService interfaces.IWalletService
}

func NewWalletController(IWalletService interfaces.IWalletService) *WalletController {
	return &WalletController{
		IWalletService: IWalletService,
	}
}

func (w *WalletController) GetWalletByUserId(userId string) *response.Response {
	wallet, err := w.IWalletService.GetWalletByUserId(userId)
	if err != nil {
		return response.BadRequest(err.Error())
	}
	return response.OK_WITH_DATA(wallet)
}
