package users

import (
	"core/pkg/errors"
	"regexp"

	"github.com/quessapp/toolkit/regexes"

	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/quessapp/toolkit/validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

// UpdateProfileDTO is DTO for payload for update user profile handler.
type UpdateProfileDTO struct {
	Nick   string `json:"nick,omitempty" bson:"nick"`
	Name   string `json:"name,omitempty" bson:"name"`
	Locale string `json:"locale,omitempty" bson:"locale"`
	Email  string `json:"email,omitempty" bson:"email"`
}

// UpdatePreferencesDTO is DTO for payload for update preferences handler.
type UpdatePreferencesDTO struct {
	EnableAPPPushNotifications bool `json:"enableAppPushNotifications" bson:"enableAppPushNotifications"`
	EnableAPPEmails            bool `json:"enableAppEmails" bson:"enableAppEmails"`
}

// Validate is a method of UpdateProfileDTO that validates the fields of the struct.
// The method uses the validation package to validate the Nick, Name, Email and Locale fields.
// The Nick, Name and Email fields are required and must have a length between 3 and 50 characters for the Nick and Name fields
// and between 5 and 200 characters for the Email field.
// The Email field must also match a valid email format using the is.Email method.
// The Locale field is required and must match the regular expression defined in the regexes package for valid locales.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d UpdateProfileDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&d.Name, validation.Required.Error(errors.NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NAME_FIELD_LENGTH)),
		validation.Field(&d.Email, validation.Required.Error(errors.EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(errors.EMAIL_FIELD_LENGTH), is.Email.Error(errors.EMAIL_FORMAT_INVALID)),
		validation.Field(&d.Locale, validation.Required.Error(errors.LOCALE_FIELD_REQUIRED), validation.Match(regexp.MustCompile(regexes.LOCALES)).Error(errors.LOCALE_FIELD_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate is a method of UpdatePreferencesDTO that validates the fields of the struct.
// The method uses the validation package to validate the EnableAPPEmails and EnableAPPPushNotifications fields.
// Both fields are required and must be present.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d UpdatePreferencesDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.EnableAPPEmails, validation.Required.Error(errors.ENABLE_APP_EMAILS_FIELD_REQUIRED)),
		validation.Field(&d.EnableAPPPushNotifications, validation.Required.Error(errors.ENABLE_APP_NOTIFICATIONS_FIELD_REQUIRED)),
	)

	return validations.GetValidationError(validationResult)
}
