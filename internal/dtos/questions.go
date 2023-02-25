package dtos

import (
	"core/internal/errors"
	"time"

	"github.com/kuriozapp/toolkit/validations"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateQuestionDTO is DTO for payload for create question handler.
type CreateQuestionDTO struct {
	ID          toolkitEntities.ID
	Content     string
	SendTo      toolkitEntities.ID
	SentBy      toolkitEntities.ID
	IsAnonymous bool
	CreatedAt   time.Time
}

// Validate validates passed struct then returns a string.
func (d CreateQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(3, 250).Error(errors.CONTENT_LENGTH)),
		validation.Field(&d.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED), validation.Length(3, 50).Error(errors.SEND_TO_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
