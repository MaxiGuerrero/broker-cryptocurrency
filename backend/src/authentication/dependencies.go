package authentication

import (
	"backend/src/system/database"
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
				authRepository: &AuthRepository{database: db},
				encrypter:      &Encrypter{},
				jwtBuilder:     NewJWTBuilder(),
			},
		},
	}
}
