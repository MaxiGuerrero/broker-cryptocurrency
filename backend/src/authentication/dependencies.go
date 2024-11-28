package authentication

type ContainerDependency struct {
	authController *AuthController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency() {
	container = &ContainerDependency{
		authController: &AuthController{
			authService: &AuthService{
				authRepository: &AuthRepository{},
			},
		},
	}
}
