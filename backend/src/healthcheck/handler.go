package healthcheck

import (
	responses "broker-cryptocurrency/backend/src/system/server/responses"

	"github.com/gofiber/fiber/v2"
)

type Healthcheck struct{}

// Entrypoint of feature
func (h *Healthcheck) Entrypoint(router fiber.Router) {
	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		response := h.controller()
		return c.Status(response.Code).JSON(response)
	})
}

func (h *Healthcheck) controller() *responses.Response {
	return responses.OK()
}
