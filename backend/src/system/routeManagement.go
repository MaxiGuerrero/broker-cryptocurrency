package system

import "github.com/gofiber/fiber/v2"

type Route interface {
	Handler(router fiber.Router)
}

func RegisterRoutes(routes *[]Route, router fiber.Router) {
	for _, route := range *routes {
		route.Handler(router)
	}
}
