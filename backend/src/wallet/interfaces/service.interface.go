package interfaces

import "backend/src/common/models"

type IWalletService interface {
	GetWalletByUserId(userId string) (*models.Wallet, error)
}
