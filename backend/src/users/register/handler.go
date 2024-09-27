package register

import (
	response "broker-cryptocurrency/backend/src/system/server/responses"
	"broker-cryptocurrency/backend/src/users/models"

	"github.com/gofiber/fiber/v2"
)

type Register struct{}

func (r *Register) Entrypoint(router fiber.Router) {
	router.Post("/register", func(c *fiber.Ctx) error {
		req := models.CreateUserRequest{}
		if parseError := c.BodyParser(&req); parseError != nil {
			return c.Status(400).JSON(response.BadRequest(parseError.Error()))
		}
		response := CreateUserController(req)
		return c.Status(response.Code).JSON(response)
	})
}
