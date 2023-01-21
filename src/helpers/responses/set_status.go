package helpers

import "github.com/gofiber/fiber/v2"

// SetResponseStatus sets response status and return it.
func SetResponseStatus(c *fiber.Ctx, status int) int {
	var currentStatus int = 200

	if status != 0 {
		currentStatus = status
	}

	c.SendStatus(currentStatus)

	return currentStatus
}
