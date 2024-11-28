package authentication

import (
	"backend/src/authentication/models"

	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Post("/register", func(c *fiber.Ctx) error {
		authController := container.authController
		req := models.RegisterRequest{}
		c.BodyParser(&req)
		response := authController.register(&req)
		return c.Status(response.Code).JSON(response)
	})
}
