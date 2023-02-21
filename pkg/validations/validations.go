package validations

import (
	"errors"
	"fmt"
	"strings"
)

func GetValidationError(validationErr error) error {
	if validationErr == nil {
		return nil
	}

	firstErrMsg := strings.Split(fmt.Sprint(validationErr), "; ")[0]
	removedErrorSuffix := strings.Split(firstErrMsg, ": ")[1]

	return errors.New(removedErrorSuffix)
}
