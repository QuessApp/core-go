package helpers

import (
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// ParseResponse returns a json with request info like status, message, data, etc.
func ParseResponse(c *fiber.Ctx, payload models.Response) error {
	return c.JSON(fiber.Map{
		"ok":      payload.Ok,
		"message": payload.Message,
		"data":    payload.Data,
		"status":  payload.Status,
	})
}

// ParseResponse returns a json with request info like status, message, data, etc.
func ParseSuccessResponse(c *fiber.Ctx, payload models.Response) error {
	return ParseResponse(c, models.Response{
		Ok:      true,
		Message: nil,
		Data:    payload.Data,
		Status:  payload.Status,
	})
}
