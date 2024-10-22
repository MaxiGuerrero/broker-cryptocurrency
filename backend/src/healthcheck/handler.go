package healthcheck

import (
	responses "broker-cryptocurrency/backend/src/system/responses"

	"github.com/gofiber/fiber/v2"
)

// Entrypoint of feature
func Handler(router fiber.Router) {
	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		response := controller()
		return c.Status(response.Code).JSON(response)
	})
}

func controller() *responses.Response {
	return responses.OK()
}
