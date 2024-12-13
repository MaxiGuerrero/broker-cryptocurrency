package middlewares

import (
	"backend/src/system"
	response "backend/src/system/responses"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const bearerPrefix = "Bearer "

func AuthGuard(c *fiber.Ctx) error {
	jwtBuilder := system.NewJWTBuilder()
	authHeader := c.Get("Authorization")
	token := strings.TrimPrefix(authHeader, bearerPrefix)
	payload, err := jwtBuilder.ValidateToken(token)
	if err != nil {
		response := response.Unauthorized()
		return c.Status(response.Code).JSON(response)
	}
	c.Locals("user", payload)
	return c.Next()
}
