package helpers

import (
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// ParseResponse returns a JSON with response info like status, message, data, etc.
func ParseResponse(c *fiber.Ctx, payload models.Response) error {
	status := SetResponseStatus(c, payload.Status)

	return c.JSON(fiber.Map{
		"ok":      payload.Ok,
		"message": payload.Message,
		"data":    payload.Data,
		"status":  status,
	})
}
