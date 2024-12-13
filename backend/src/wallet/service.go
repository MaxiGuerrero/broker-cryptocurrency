package wallet

import (
	"backend/src/common/interfaces"
	"backend/src/common/models"
	"errors"
)

type WalletService struct {
	iWalletRepository interfaces.IWalletRepository
}

func NewWalletService(iWalletRepository interfaces.IWalletRepository) *WalletService {
	return &WalletService{
		iWalletRepository: iWalletRepository,
	}
}

func (w *WalletService) GetWalletByUserId(userId string) (*models.Wallet, error) {
	wallet := w.iWalletRepository.GetWalletByUserId(userId)
	if wallet == nil {
		return nil, errors.New("wallet not created, contact to administration")
	}
	return wallet, nil
}
