package interfaces

import "backend/src/common/models"

type IWalletRepository interface {
	CreateWallet(userId string)
	GetWalletByUserId(userId string) *models.Wallet
}
