package wallet

import (
	"backend/src/system"
	"backend/src/system/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	walletController := container.walletController
	router.Get("/funds", middlewares.AuthGuard, func(c *fiber.Ctx) error {
		user := c.Locals("user").(system.Payload)
		response := walletController.GetWalletByUserId(user.UserId)
		return c.Status(response.Code).JSON(response)
	})
}
