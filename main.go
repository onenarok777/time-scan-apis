package main

import (
	"log"
	"time-scan/internal/routes"
	"time-scan/pkg/gorm"
	"time-scan/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// init utils
	utils.InitDotEnv()

	// init db
	gorm.InitDB()
	gorm.Migrate(gorm.DB)

	// init fiber
	app := fiber.New()

	// use cors
	app.Use(cors.New())

	// enable log api
	app.Use(logger.New())

	// set route
	routes.SetRoutes(app)

	log.Fatal(app.Listen(":4001"))
}
