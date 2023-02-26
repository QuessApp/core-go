package blocks

import (
	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"github.com/kuriozapp/toolkit/validations"

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
func (d BlockUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.UserToBlock, validation.Required.Error(USER_TO_BLOCK_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
