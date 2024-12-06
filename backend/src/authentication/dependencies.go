package authentication

import "backend/src/system"

type ContainerDependency struct {
	authController *AuthController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency(db *system.Database) {
	container = &ContainerDependency{
		authController: &AuthController{
			authService: &AuthService{
				authRepository: &AuthRepository{database: db},
				encrypter:      &system.Encrypter{},
			},
		},
	}
}
