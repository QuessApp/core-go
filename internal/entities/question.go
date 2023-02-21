package entities

import (
	"core/pkg/entities"
	"core/pkg/errors"
	"core/pkg/validations"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Questions is a model for each question in app.
type Question struct {
	ID      entities.ID `json:"id" bson:"_id"`
	Content string      `json:"content"`

	SendTo string `json:"sendTo" bson:"sendTo"`
	SentBy string `json:"sentBy,omitempty" bson:"sentBy"`
	// Reply is replied data content. Type must be Entities.Reply or nil.
	Reply any `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous" bson:"IsAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty" bson:"isHiddenByReceiver"`
	IsReplied          bool `json:"isReplied" bson:"isReplied"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// Validate validates passed struct then returns a string.
func (q Question) Validate() error {
	validationResult := validation.ValidateStruct(&q,
		validation.Field(&q.Content, validation.Required.Error(errors.CONTENT_REQUIRED), validation.Length(3, 250).Error(errors.CONTENT_LENGTH)),
		validation.Field(&q.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED), validation.Length(3, 50).Error(errors.SEND_TO_LENGTH)),
		validation.Field(&q.IsAnonymous, validation.Required.Error(errors.IS_ANONYMOUS_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
