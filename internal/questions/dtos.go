package questions

import (
	"time"

	"github.com/quessapp/toolkit/validations"

	"github.com/quessapp/core-go/pkg/errors"

	toolkitEntities "github.com/quessapp/toolkit/entities"

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

// ReplyQuestionDTO is DTO for payload for reply question handler.
type ReplyQuestionDTO struct {
	ID      toolkitEntities.ID
	Content string
}

// EditQuestionReplyDTO is DTO for payload for edit reply question handler.
type EditQuestionReplyDTO struct {
	ID                  toolkitEntities.ID
	Content             string
	OldContent          string
	OldContentCreatedAt time.Time
}

// Validate is a method of ReplyQuestionDTO that validates the fields of the struct.
// The method uses the validation package to validate the Content field.
// The Content field is required and must have a length between 1 and 250 characters.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d ReplyQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate is a method of EditQuestionReplyDTO that validates the fields of the struct.
// The method uses the validation package to validate the Content field.
// The Content field is required and must have a length between 1 and 250 characters.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d EditQuestionReplyDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate is a method of CreateQuestionDTO that validates the fields of the struct.
// The method uses the validation package to validate the Content and SendTo fields.
// The Content field is required and must have a length between 1 and 250 characters.
// The SendTo field is required and must have a length between 3 and 50 characters.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d CreateQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
		validation.Field(&d.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED), validation.Length(3, 50).Error(errors.SEND_TO_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
