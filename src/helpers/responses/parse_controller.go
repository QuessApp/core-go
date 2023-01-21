package helpers

import (
	"core/src/models"

	"github.com/gofiber/fiber/v2"
)

// ParseControllerResponse will handle controller requests.
// If controller returns an error, the above function will format the error.
// Otherwise, it returns the success response with data.
func ParseControllerResponse(c *fiber.Ctx, err models.RequestError, data interface{}) error {
	if err.Message != nil {
		return ParseResponse(c, models.Response{
			Ok:      false,
			Message: err.Message,
			Data:    nil,
			Status:  err.Status,
		})
	}

	return ParseSuccessResponse(c, models.Response{
		Status: 201,
		Data:   data,
	})
}
