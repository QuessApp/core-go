package questions

import (
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

// ReplyQuestionDTO is DTO for payload for reply question handler.
type ReplyQuestionDTO struct {
	ID      toolkitEntities.ID
	Content string
}

// Validate validates passed struct then returns a string
//
// It validates if content is valid.
func (d ReplyQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(CONTENT_REQUIRED), validation.Length(1, 250).Error(CONTENT_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate validates passed struct then returns a string.
//
// It validates if content and send to fields are valids.
func (d CreateQuestionDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Content, validation.Required.Error(CONTENT_REQUIRED), validation.Length(1, 250).Error(CONTENT_LENGTH)),
		validation.Field(&d.SendTo, validation.Required.Error(SEND_TO_REQUIRED), validation.Length(3, 50).Error(SEND_TO_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
