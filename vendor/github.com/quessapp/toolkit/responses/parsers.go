package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/errors"
)

// ParseSuccessful parses a successfull response and normalizes to a specif json format.
func ParseSuccessful(c *fiber.Ctx, status int, data any) error {
	res := &entities.Response{
		Ok:      true,
		Error:   false,
		Message: nil,
		Data:    data,
	}

	c.Status(status)
	return c.JSON(res)
}

// ParseSuccessful parses a unsuccesfull response and normalizes to a specif json format.
func ParseUnsuccesfull(c *fiber.Ctx, status int, err string) error {
	res := &entities.Response{
		Ok:      false,
		Error:   true,
		Message: err,
		Data:    nil,
	}

	c.Status(status)
	return c.JSON(res)
}

// HasRecordsInMongo returns a bool value if message error is equal to the Mongo no result documents error.
func HasRecordsInMongo(err error) bool {
	return err.Error() == errors.MONGO_NO_DOCUMENTS
}
