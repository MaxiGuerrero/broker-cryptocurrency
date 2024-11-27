package authentication

type ContainerDependency struct {
	authController *AuthController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency() {
	container = &ContainerDependency{
		authController: &AuthController{
			AuthService: &AuthService{
				AuthRepository: &AuthRepository{},
			},
		},
	}
}
