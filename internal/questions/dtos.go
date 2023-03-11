package questions

import (
	"time"

	"github.com/quessapp/toolkit/validations"

	"core/pkg/errors"

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

// Validate validates passed struct then returns a string
//
// It validates if content is valid.
func (d ReplyQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate validates passed struct then returns a string
//
// It validates if content is valid.
func (d EditQuestionReplyDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate validates passed struct then returns a string.
//
// It validates if content and send to fields are valids.
func (d CreateQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(1, 250).Error(errors.CONTENT_LENGTH)),
		validation.Field(&d.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED), validation.Length(3, 50).Error(errors.SEND_TO_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
