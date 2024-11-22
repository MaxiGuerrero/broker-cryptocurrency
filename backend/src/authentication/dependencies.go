package authentication

type ContainerDependency struct {
	healthCheckController *HealthCheckController
}

var container *ContainerDependency

// Initialize container dependency
func InitContainerDependency() {
	container = &ContainerDependency{
		healthCheckController: &HealthCheckController{},
	}
}
