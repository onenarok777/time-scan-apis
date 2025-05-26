package handlers

import (
	"time-scan/internal/auth/models/dtx"
	"time-scan/internal/auth/services"

	"github.com/gofiber/fiber/v2"
)

func LoginHandle(c *fiber.Ctx) error {
	// body parser
	var req dtx.RequestLogin
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtx.ErrorResponse{
			Error: "Body Not Found !",
		})

	}

	// call service
	res, err := services.LoginService(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtx.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
