package dtos

import (
	pkg "core/pkg/entities"
	"core/pkg/errors"
	"core/pkg/validations"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateQuestionDTO is DTO for payload for create question handler.
type CreateQuestionDTO struct {
	ID          pkg.ID
	Content     string
	SendTo      pkg.ID
	SentBy      pkg.ID
	IsAnonymous bool
	CreatedAt   time.Time
}

// Validate validates passed struct then returns a string.
func (q CreateQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&q,
		validation.Field(&q.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(3, 250).Error(errors.CONTENT_LENGTH)),
		validation.Field(&q.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED), validation.Length(3, 50).Error(errors.SEND_TO_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
