package healthcheck

import "sync"

// Dependency container pattern that will be used along to the package.
type ContainerDependency struct {
	healthCheckController *HealthCheckController
}

var container *ContainerDependency
var once sync.Once

func InitContainer() {
	once.Do(func() {
		container = &ContainerDependency{
			healthCheckController: &HealthCheckController{},
		}
	})
}

func getDependencyContainer() *ContainerDependency {
	if container == nil {
		panic("Container has not been initialized")
	}
	return container
}
