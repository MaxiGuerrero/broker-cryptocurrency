package system

import "github.com/gofiber/fiber/v2"

type Handler func(router fiber.Router)

func PlugInFeatures(handlers *[]Handler, router fiber.Router) {
	for _, handler := range *handlers {
		handler(router)
	}
}
