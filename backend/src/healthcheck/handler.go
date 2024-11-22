package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

// Entrypoint of feature
func Handler(router fiber.Router) {
	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		healthCheckController := container.healthCheckController
		response := healthCheckController.controller()
		return c.Status(response.Code).JSON(response)
	})
}
