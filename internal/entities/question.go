package entities

import (
	pkg "core/pkg/entities"
	"core/pkg/errors"
	"core/pkg/validations"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Questions is a model for each question in app.
type Question struct {
	ID      pkg.ID `json:"id" bson:"_id,omitempty"`
	Content string `json:"content"`

	// SendTo represents the user that will receive the question. Type must bet Entities.ID, nill ou Entities.User
	SendTo any `json:"sendTo,omitempty" bson:"sendTo"`
	// SentBy represents the user who sent the question. Type must bet Entities.ID, nill ou Entities.User
	SentBy any `json:"sentBy,omitempty" bson:"sentBy"`
	// Reply is replied data content. Type must be Entities.Reply or nil.
	Reply any `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous" bson:"IsAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty" bson:"isHiddenByReceiver"`
	IsReplied          bool `json:"isReplied,omitempty" bson:"isReplied"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
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

// MapAnonymousFields maps question in order to hide who sent the question, if the questions is anonymous.Otherwise, just returns the whole data.
func (q Question) MapAnonymousFields() *Question {
	if q.IsAnonymous {
		return &Question{
			ID:          q.ID,
			Content:     q.Content,
			IsAnonymous: q.IsAnonymous,
			CreatedAt:   q.CreatedAt,
		}
	}
	return &q
}
