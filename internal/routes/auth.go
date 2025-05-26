package routes

import (
	"time-scan/internal/auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", handlers.LoginHandle)
	auth.Get("/me")
}
