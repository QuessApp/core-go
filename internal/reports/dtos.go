package reports

import (
	"core/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/quessapp/toolkit/validations"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// CreateReportDTO is DTO for payload for create report handler.
type CreateReportDTO struct {
	ID     toolkitEntities.ID `json:"id" bson:"_id"`
	Type   string             `json:"type" bson:"type"`
	Reason string             `json:"reason" bson:"reason"`
	SendTo toolkitEntities.ID `json:"sendTo" bson:"sendTo"`
	SentBy toolkitEntities.ID `json:"sentBy" bson:"sentBy"`
}

// Validate function is responsible for validating the CreateReportDTO struct.
//
// It receives a CreateReportDTO struct and validates its fields:
// - Reason: must be one of the allowed values (spam, nudity or sexual activity, hate speech or symbols,
// violence or dangerous organizations, bullying or harassment, selling illegal or regulated goods,
// intellectual property violations, suicide or self-injury, eating disorders, scams or fraud, false information).
// - Type: must be one of the allowed values (question, user).
// - SendTo: must not be empty.
//
// If any of the fields fail validation, an error is returned.
// Otherwise, nil is returned.
func (d CreateReportDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Reason, validation.Required.Error(errors.REASON_FIELD_REQUIRED), validation.In(
			"spam",
			"nudity or sexual activity",
			"hate speech or symbols",
			"violence or dangerous organizations",
			"bullying or harassment",
			"selling illegal or regulated goods",
			"intellectual property violations",
			"suicide or self-injury",
			"eating disorders",
			"scams or fraud",
			"false information",
		).Error(errors.REASON_FIELD_INVALID)),
		validation.Field(&d.Type, validation.Required.Error(errors.TYPE_FIELD_REQUIRED), validation.In("question", "user").Error(errors.TYPE_FIELD_INVALID)),
		validation.Field(&d.SendTo, validation.Required.Error(errors.SEND_TO_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
