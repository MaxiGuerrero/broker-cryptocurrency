package authentication

import (
	"backend/src/system"
	"backend/src/system/database"
	"backend/src/wallet"
)

type ContainerDependency struct {
	authController *AuthController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency(db *database.Database) {
	container = &ContainerDependency{
		authController: &AuthController{
			authService: &AuthService{
				authRepository: &AuthRepository{database: db, iWalletRepository: &wallet.WalletRepository{Database: db}},
				encrypter:      &Encrypter{},
				jwtBuilder:     system.NewJWTBuilder(),
			},
		},
	}
}
