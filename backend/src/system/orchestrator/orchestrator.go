package orchestrator

import "github.com/gofiber/fiber/v2"

type Initializer interface {
	Entrypoint(router fiber.Router)
}

func Orchestration(initializers *[]Initializer, router fiber.Router) {
	for _, initializer := range *initializers {
		initializer.Entrypoint(router)
	}
}
