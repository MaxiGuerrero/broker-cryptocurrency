package authentication

import (
	"backend/src/authentication/models"

	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	authController := container.authController
	router.Post("/register", func(c *fiber.Ctx) error {
		req := models.RegisterRequest{}
		c.BodyParser(&req)
		response := authController.Register(&req)
		return c.Status(response.Code).JSON(response)
	})
	router.Post("/login", func(c *fiber.Ctx) error {
		req := models.LoginRequest{}
		c.BodyParser(&req)
		response := authController.Login(&req)
		return c.Status(response.Code).JSON(response)
	})
}
