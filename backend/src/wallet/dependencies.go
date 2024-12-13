package wallet

import (
	"backend/src/system/database"
)

type ContainerDependency struct {
	walletController WalletController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency(db *database.Database) {
	container = &ContainerDependency{
		walletController: WalletController{
			IWalletService: &WalletService{
				iWalletRepository: &WalletRepository{
					Database: db,
				},
			},
		},
	}
}
