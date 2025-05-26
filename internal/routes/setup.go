package routes

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")

	RouteAuth(api)
	RouteTimeScan(api)
}
