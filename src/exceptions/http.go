package exceptions

import (
	"core/src/helpers"
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// NewHttpException returns a json with request info like status, message, data, etc.
func NewHttpException(c *fiber.Ctx, payload models.Response) error {
	c.SendStatus(payload.Status)

	return helpers.ParseResponse(c, models.Response{
		Ok:      false,
		Message: payload.Message,
		Data:    nil,
		Status:  payload.Status,
	})
}
