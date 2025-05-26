package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadTime(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(file)
	return nil
}
