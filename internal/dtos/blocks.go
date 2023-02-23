package dtos

import (
	pkg "core/pkg/entities"
	"core/pkg/errors"
	"core/pkg/validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

// BlockUserDTO is a dto for payload to block a user.
type BlockUserDTO struct {
	ID          pkg.ID `json:"id" bson:"_id" `
	UserToBlock pkg.ID `json:"userToBlock" bson:"userToBlock"`
	BlockedBy   pkg.ID `json:"blockedBy" bson:"blockedBy"`
}

// Validate validates passed struct then returns a string.
func (d BlockUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.UserToBlock, validation.Required.Error(errors.USER_TO_BLOCK_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
