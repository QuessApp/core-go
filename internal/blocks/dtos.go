package blocks

import (
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/quessapp/toolkit/validations"

	"core/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

// BlockUserDTO is a dto for payload to block an user.
type BlockUserDTO struct {
	ID          toolkitEntities.ID `json:"id" bson:"_id"`
	UserToBlock toolkitEntities.ID `json:"userToBlock" bson:"userToBlock"`
	BlockedBy   toolkitEntities.ID `json:"blockedBy" bson:"blockedBy"`
}

// UnblockUserDTO is a dto for payload to unblock an user.
type UnblockUserDTO struct {
	BlockedUserID toolkitEntities.ID
}

// Validate validates passed struct then returns a string.
//
// It validates if user to block id is valid.
func (d BlockUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.UserToBlock, validation.Required.Error(errors.USER_TO_BLOCK_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
