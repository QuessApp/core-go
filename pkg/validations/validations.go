package validations

import (
	"errors"
	"fmt"
	"strings"
)

// GetValidationError receives an error from validator lib and parses error getting the first error instead of an object.
func GetValidationError(validationErr error) error {
	if validationErr == nil {
		return nil
	}

	firstErrMsg := strings.Split(fmt.Sprint(validationErr), "; ")[0]
	removedErrorSuffix := strings.Split(firstErrMsg, ": ")[1]

	return errors.New(removedErrorSuffix)
}
