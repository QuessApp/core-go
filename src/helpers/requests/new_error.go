package helpers

import (
	"core/src/models"
	"fmt"
)

// NewRequestError just receives an error and parses it to string and to RequestError model.
// This function help us to show errors.
func NewRequestError(err error, status int) models.RequestError {
	return models.RequestError{
		Message: fmt.Sprint(err),
		Status:  status,
	}
}
