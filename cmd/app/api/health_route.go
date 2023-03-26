package api

import (
	"placio-app/controller"
	"placio-app/utility"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(api fiber.Router) {
	api.Get("/health", utility.Use(controller.HealthCheck))

}
