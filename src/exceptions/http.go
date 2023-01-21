package exceptions

import (
	helpers "core/src/helpers/responses"
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// HttpException returns a JSON with request info like status, message, data, etc.
func HttpException(c *fiber.Ctx, payload models.Response) error {
	c.SendStatus(payload.Status)

	return helpers.ParseResponse(c, models.Response{
		Ok:      false,
		Message: payload.Message,
		Data:    nil,
		Status:  payload.Status,
	})
}
