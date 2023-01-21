package helpers

import (
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// ParseResponse returns a JSON with response info like status, message, data, etc.
func ParseSuccessResponse(c *fiber.Ctx, payload models.Response) error {
	status := SetResponseStatus(c, payload.Status)

	return ParseResponse(c, models.Response{
		Ok:      true,
		Message: nil,
		Data:    payload.Data,
		Status:  status,
	})
}
