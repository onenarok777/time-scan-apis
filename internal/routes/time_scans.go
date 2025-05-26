package routes

import (
	"time-scan/internal/time_scans/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteTimeScan(api fiber.Router) {
	timeScan := api.Group("/time-scan")

	timeScan.Post("/upload", handlers.UploadTime)
}
