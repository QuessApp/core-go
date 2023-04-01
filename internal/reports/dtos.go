package reports

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"
	"github.com/quessapp/core-go/pkg/reports"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/validations"
)

// CreateReportDTO is DTO for payload for create report handler.
type CreateReportDTO struct {
	ID     toolkitEntities.ID `json:"id" bson:"_id"`
	Type   string             `json:"type" bson:"type"`
	Reason string             `json:"reason" bson:"reason"`
	SendTo toolkitEntities.ID `json:"sendTo" bson:"sendTo"`
	SentBy toolkitEntities.ID `json:"sentBy" bson:"sentBy"`
}

// checkIfReasonIsValid is a function that checks if the given reason is valid.
// It receives a reason and checks if it is one of the allowed values.
// If it is, nil is returned.
// Otherwise, an error is returned.
func checkIfReasonIsValid(value any) error {
	reasonsAsList := strings.Split(reports.REASONS, ", ")
	s, _ := value.(string)

	for _, r := range reasonsAsList {
		if r == s {
			return nil
		}
	}

	return errors.New(pkgErrors.REASON_FIELD_INVALID)
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
		validation.Field(&d.Reason, validation.Required.Error(pkgErrors.REASON_FIELD_REQUIRED), validation.By(checkIfReasonIsValid)),
		validation.Field(&d.Type, validation.Required.Error(pkgErrors.TYPE_FIELD_REQUIRED), validation.In("question", "user").Error(pkgErrors.TYPE_FIELD_INVALID)),
		validation.Field(&d.SendTo, validation.Required.Error(pkgErrors.SEND_TO_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
